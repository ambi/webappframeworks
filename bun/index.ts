import Handlebars from 'handlebars';
import mysql from 'mysql2/promise';

type World = {
  id: number;
  randomNumber: number;
};

type Fortune = {
  id: number;
  message: string;
};

// const db = await new Client().connect({
//   hostname: '127.0.0.1',
//   username: 'root',
//   db: 'hello_world',
//   password: '',
// });
const db = mysql.createPool({
  host: '127.0.0.1',
  user: 'root',
  database: 'hello_world',
});

const fortunesTmplFile = await Bun.file("fortunes.handlebars").text();
const fortunesTmpl = Handlebars.compile(fortunesTmplFile);
const responseHeaders = { 'Server': 'Bun' };

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

  return new Response(fortunesTmpl({ fortunes }), {
    headers: {
      ...responseHeaders,
      'Content-Type': 'text/html; UTF-8',
  }});
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
      return await queriesHandler(parseInt(url.searchParams.get('n') || '0'));
    case '/updates':
      return await updatesHandler(parseInt(url.searchParams.get('n') || '0'));
    case '/fortunes':
      return await fortunesHandler();
    default:
      return new Response('404 Not Found', {
        status: 404,
        headers: responseHeaders,
      });
  }
}

const server = Bun.serve({
  port: 8080,
  reusePort: true,
  fetch: handler,
});

console.log(`Listening on localhost:${server.port}`);
