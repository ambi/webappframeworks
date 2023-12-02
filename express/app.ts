import cluster from 'cluster';
import os from 'node:os';
import express from 'express';
import { engine } from 'express-handlebars';

import { db } from './database.js';

function serve() {
  const app = express();

  app.use(express.urlencoded({ extended: true }));

  // Set headers for all routes
  app.use((req, res, next) => {
    res.setHeader("Server", "Express");
    return next();
  });

  app.engine('handlebars', engine());
  app.set('view engine', 'handlebars');

  app.get('/json', (req, res) => res.send({ message: 'Hello, World!' }));

  app.get('/plaintext', (req, res) => res.header('Content-Type', 'text/plain').send('Hello, World!'));

  app.get('/db', async (req, res) => {
    const id = Math.floor(Math.random() * 10000) + 1;
    const world = await db.selectFrom('world').selectAll().where('id', '=', id).executeTakeFirstOrThrow();

    res.json(world);
  });

  app.get('/queries', async (req, res) => {
    const results = [];
    const queries = Math.min(queryNumber(req.query.n as string | string[] | undefined) || 1, 500);

    for (let i = 1; i <= queries; i++) {
      const id = Math.floor(Math.random() * 10000) + 1;
      const world = await db.selectFrom('world').selectAll().where('id', '=', id).executeTakeFirstOrThrow();
      results.push(world);
    }

    res.json(results);
  });

  app.get('/updates', async (req, res) => {
    const results = [];
    const queries = Math.min(queryNumber(req.query.n as string | string[] | undefined) || 1, 500);

    for (let i = 1; i <= queries; i++) {
      const id = Math.floor(Math.random() * 10000) + 1;
      const world = await db.selectFrom('world').selectAll().where('id', '=', id).executeTakeFirstOrThrow();

      world.randomNumber = ~~(Math.random() * 10000) + 1;
      await db.updateTable('world').set(world).where('id', '=', id).execute();
      results.push(world);
    }

    res.json(results);
  });

  app.get('/fortunes', async (req, res) => {
    const fortunes = await db.selectFrom('fortune').selectAll().execute();

    const newFortune = { id: 0, message: 'Additional fortune added at request time.' };
    fortunes.push(newFortune);
    fortunes.sort((a, b) => (a.message < b.message ? -1 : 1));

    res.render('fortunes', { fortunes: fortunes });
  });

  app.listen(8080);
}

function queryNumber(value: string | string[] | undefined) {
  if (!value) return 0;
  if (typeof value === 'string') {
    return parseInt(value);
  }
  if (value.length === 0) return 0;
  return parseInt(value[0]);
}

const numCPUs = os.cpus().length;

if (cluster.isPrimary) {
  console.log(`Primary ${process.pid} is running`);

  for (let i = 0; i < numCPUs; i++) {
    cluster.fork();
  }

  cluster.on('exit', (worker, code, signal) => {
    console.log(`worker ${worker.process.pid} died`);
  });
} else {
  console.log(`Worker ${process.pid} is running`);
  serve();
}

// serve();