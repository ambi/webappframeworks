import { Kysely, MysqlDialect } from 'kysely';
import { createPool } from 'mysql2';

import type { Database } from './types.js';

const dialect = new MysqlDialect({
  pool: createPool({
    database: 'hello_world',
    host: 'localhost',
    user: 'root',
    password: '',
    // connectionLimit: 10,
  }),
});

export const db = new Kysely<Database>({ dialect });
