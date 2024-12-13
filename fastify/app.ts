import cluster from 'node:cluster';
import os from 'node:os';
import { parseArgs }  from 'node:util';
import view from '@fastify/view';
import Fastify from 'fastify';
import handlebars from 'handlebars';

import { db } from './database.js';

async function serve() {
  const app = Fastify();

  app.register(view, {
    engine: { handlebars },
    viewExt: 'handlebars',
    layout: 'views/layouts/main.handlebars',
  });

  app.get('/json', (req, res) => {
    return res.send({ message: 'Hello, World!' });
  });

  app.get('/plaintext', (req, res) => {
    return res.header('Content-Type', 'text/plain').send('Hello, World!');
  });

  app.get('/db', async (req, res) => {
    const id = Math.floor(Math.random() * 10000) + 1;
    const world = await db.selectFrom('world').selectAll().where('id', '=', id).executeTakeFirstOrThrow();

    return res.send(world);
  });

  app.get<{ Querystring: { n: string } }>('/queries', async (req, res) => {
    const results = [];
    const queries = Math.min(queryNumber(req.query.n) || 1, 500);

    for (let i = 1; i <= queries; i++) {
      const id = Math.floor(Math.random() * 10000) + 1;
      const world = await db.selectFrom('world').selectAll().where('id', '=', id).executeTakeFirstOrThrow();
      results.push(world);
    }

    return res.send(results);
  });

  app.get<{ Querystring: { n: string } }>('/updates', async (req, res) => {
    const results = [];
    const queries = Math.min(queryNumber(req.query.n) || 1, 500);

    for (let i = 1; i <= queries; i++) {
      const id = Math.floor(Math.random() * 10000) + 1;
      const world = await db.selectFrom('world').selectAll().where('id', '=', id).executeTakeFirstOrThrow();

      world.randomNumber = ~~(Math.random() * 10000) + 1;
      await db.updateTable('world').set(world).where('id', '=', id).execute();
      results.push(world);
    }

    return res.send(results);
  });

  app.get('/fortunes', async (req, res) => {
    const fortunes = await db.selectFrom('fortune').selectAll().execute();

    const newFortune = { id: 0, message: 'Additional fortune added at request time.' };
    fortunes.push(newFortune);
    fortunes.sort((a, b) => (a.message < b.message ? -1 : 1));

    return res.view('views/fortunes', { fortunes: fortunes });
  });

  await app.listen({ port: 8080 });
}

function queryNumber(value: string | string[] | undefined) {
  if (!value) return 0;
  if (typeof value === 'string') {
    return Number.parseInt(value);
  }
  if (value.length === 0) return 0;
  return Number.parseInt(value[0]);
}

function runClusterMode() {
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
    console.log(`Secondary ${process.pid} is running`);

    serve().catch(console.error);
  }
}

function main() {
  const options = {
    cluster: {
      type: 'boolean',
      short: 'c',
      default: false
    },
  } as const;
  const { values } = parseArgs({ options });

  if (values.cluster) {
    runClusterMode();
  } else {
    serve().catch(console.error);
  }
}
main();