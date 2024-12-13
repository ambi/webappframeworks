import mysql from 'npm:mysql2/promise';

export type World = {
  id: number;
  randomNumber: number;
};

export type Fortune = {
  id: number;
  message: string;
};

export const db = mysql.createPool({
  host: 'localhost',
  user: 'root',
  database: 'hello_world',
});
