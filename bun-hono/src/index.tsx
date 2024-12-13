/** @jsx jsx */
/** @jsxImportSource hono/jsx */

import { type Context, Hono } from 'hono';
import { getConnInfo } from 'hono/bun';

import { Fortunes } from './Fortunes';
import { type World, type Fortune, db } from './database';

function setResponseHeader(c: Context) {
  c.header('Server', 'Bun');
}

function getN(c: Context) {
  const n = c.req.query('n');
  return n ? Number.parseInt(n) : 0;
}

function jsonHandler(c: Context) {
  setResponseHeader(c);

  return c.json({ message: 'Hello, World!' });
}

function plaintextHandler(c: Context) {
  setResponseHeader(c);

  return c.text('Hello, World!');
}

async function dbHandler(c: Context) {
  setResponseHeader(c);

  const id = Math.floor(Math.random() * 10000) + 1;
  const [rows] = await db.query('SELECT * FROM world WHERE id=?', [id]);
  const worlds = rows as World[];

  return c.json(worlds[0]);
}

async function queriesHandler(c: Context) {
  setResponseHeader(c);
  const n = getN(c);

  const results: World[] = [];

  for (let i = 1; i <= n; i++) {
    const id = Math.floor(Math.random() * 10000) + 1;
    const [rows] = await db.query('SELECT * FROM world WHERE id=?', [id]);
    const worlds = rows as World[];
    results.push(worlds[0]);
  }

  return c.json(results);
}

async function updatesHandler(c: Context) {
  setResponseHeader(c);
  const n = getN(c);

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

  return c.json(results);
}

async function fortunesHandler(c: Context) {
  setResponseHeader(c);

  const [rows] = await db.query('SELECT * FROM fortune');
  const fortunes = rows as Fortune[];

  const newFortune = { id: 0, message: 'Additional fortune added at request time.' };
  fortunes.push(newFortune);
  fortunes.sort((a, b) => (a.message < b.message ? -1 : 1));

  return c.html(<Fortunes fortunes={fortunes} />);
}

const app = new Hono();

app.get('/json', jsonHandler);
app.get('/plaintext', plaintextHandler);
app.get('/db', dbHandler);
app.get('/queries', queriesHandler);
app.get('/updates', updatesHandler);
app.get('/fortunes', fortunesHandler);

export default { port: 8080, fetch: app.fetch };
