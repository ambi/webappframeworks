// deno run --allow-net index.tsx

/** @jsx jsx */
/** @jsxImportSource hono/jsx */
import { type Context, Hono } from 'https://deno.land/x/hono@v4.3.11/mod.ts';

import { Fortunes } from './Fortunes.tsx';
import { type Fortune, type World, db } from './database.ts';

function jsonHandler(c: Context) {
  return c.json({ message: 'Hello, World!' });
}

function plaintextHandler(c: Context) {
  return c.text('Hello, World!');
}

async function dbHandler(c: Context) {
  const id = Math.floor(Math.random() * 10000) + 1;
  const [rows] = await db.query('SELECT * FROM world WHERE id=?', [id]);
  const worlds = rows as World[];

  return c.json(worlds[0]);
}

async function queriesHandler(c: Context) {
  const results = [];
  const n = validateNParam(c.req.query('n'));

  for (let i = 1; i <= n; i++) {
    const id = Math.floor(Math.random() * 10000) + 1;
    const [rows] = await db.query('SELECT * FROM world WHERE id=?', [id]);
    const worlds = rows as World[];
    results.push(worlds[0]);
  }

  return c.json(results);
}

async function updatesHandler(c: Context) {
  const results = [];
  const n = validateNParam(c.req.query('n'));

  for (let i = 1; i <= n; i++) {
    const id = Math.floor(Math.random() * 10000) + 1;
    const [rows] = await db.query('SELECT * FROM world WHERE id=?', [id]);
    const worlds = rows as World[];
    const world = worlds[0];

    world.randomNumber = ~~(Math.random() * 10000) + 1;
    await db.execute('UPDATE world SET randomNumber=? WHERE id=?', [world.randomNumber, id]);
    results.push(world);
  }

  return c.json(results);
}

async function fortunesHandler(c: Context) {
  const [rows] = await db.query('SELECT * FROM fortune');
  const fortunes = rows as Fortune[];

  const newFortune = { id: 0, message: 'Additional fortune added at request time.' };
  fortunes.push(newFortune);
  fortunes.sort((a, b) => (a.message < b.message ? -1 : 1));

  return c.html(<Fortunes fortunes={fortunes} />);
}

function validateNParam(value: string | undefined) {
  if (!value) return 0;
  const n = Number.parseInt(value);
  return Math.min(n, 500);
}

const app = new Hono();
app.use('*', async (c, next) => {
  await next();
  c.res.headers.set('Server', 'Deno-Hono');
});

app.get('/json', jsonHandler);
app.get('/plaintext', plaintextHandler);
app.get('/db', dbHandler);
app.get('/queries', queriesHandler);
app.get('/updates', updatesHandler);
app.get('/fortunes', fortunesHandler);

Deno.serve({ port: 8080 }, app.fetch);
