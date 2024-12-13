// deno run --allow-net --allow-read main.tsx

import { render } from 'https://esm.sh/preact-render-to-string';

import { Fortunes } from './Fortunes.tsx';
import { type Fortune, type World, db } from './database.ts';

// import { Eta } from 'https://deno.land/x/eta@v3.0.3/src/index.ts';
// const eta = new Eta({ views: Deno.cwd(), cache: true });

const responseHeaders = {
  Server: 'Deno',
};

function jsonHandler() {
  return Response.json({ message: 'Hello, World!' }, { headers: responseHeaders });
}

function plaintextHandler() {
  return new Response('Hello, World!', { headers: responseHeaders });
}

async function dbHandler() {
  const id = Math.floor(Math.random() * 10000) + 1;
  const [rows] = await db.query('SELECT * FROM world WHERE id=?', [id]);
  const worlds = rows as World[];

  return Response.json(worlds[0], { headers: responseHeaders });
}

async function queriesHandler(n: number) {
  const results = [];

  for (let i = 1; i <= n; i++) {
    const id = Math.floor(Math.random() * 10000) + 1;
    const [rows] = await db.query('SELECT * FROM world WHERE id=?', [id]);
    const worlds = rows as World[];
    results.push(worlds[0]);
  }

  return Response.json(results, { headers: responseHeaders });
}

async function updatesHandler(n: number) {
  const results: World[] = [];

  for (let i = 1; i <= n; i++) {
    const id = Math.floor(Math.random() * 10000) + 1;
    const [rows] = await db.query('SELECT * FROM world WHERE id=?', [id]);
    const worlds = rows as World[];
    const world = worlds[0];

    world.randomNumber = ~~(Math.random() * 10000) + 1;
    await db.execute('UPDATE world SET randomNumber=? WHERE id=?', [world.randomNumber, id]);
    results.push(world);
  }

  return Response.json(results, { headers: responseHeaders });
}

async function fortunesHandler() {
  const [rows] = await db.query('SELECT * FROM fortune');
  const fortunes = rows as Fortune[];

  const newFortune = { id: 0, message: 'Additional fortune added at request time.' };
  fortunes.push(newFortune);
  fortunes.sort((a, b) => (a.message < b.message ? -1 : 1));

  return new Response(render(<Fortunes fortunes={fortunes} />), {
    headers: {
      ...responseHeaders,
      'Content-Type': 'text/html; UTF-8',
    },
  });
  // return new Response(eta.render('fortunes', { fortunes }), {
  //   headers: {
  //     ...responseHeaders,
  //     'Content-Type': 'text/html; UTF-8',
  // }});
}

async function handler(req: Request) {
  if (req.method !== 'GET') {
    return new Response('405 Method Not Allowed', {
      status: 405,
      headers: responseHeaders,
    });
  }

  const url = new URL(req.url);
  switch (url.pathname) {
    case '/json':
      return jsonHandler();
    case '/plaintext':
      return plaintextHandler();
    case '/db':
      return await dbHandler();
    case '/queries':
      return await queriesHandler(Number.parseInt(url.searchParams.get('n') || '0'));
    case '/updates':
      return await updatesHandler(Number.parseInt(url.searchParams.get('n') || '0'));
    case '/fortunes':
      return await fortunesHandler();
    default:
      return new Response('404 Not Found', {
        status: 404,
        headers: responseHeaders,
      });
  }
}

Deno.serve({
  handler,
  onError(err) {
    console.error(err);
    Deno.exit(9);
  },
  port: 8080,
  hostname: '0.0.0.0',
});
