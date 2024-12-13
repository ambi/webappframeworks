/** @jsx jsx */
/** @jsxImportSource hono/jsx */

import type { Fortune } from './database';

export function Fortunes(props: { fortunes: Fortune[] }) {
  return (
    <html>
      <head>
        <meta charset="utf-8" />
        <title>Hono App</title>
      </head>

      <body>
        <table>
          <tr>
            <th>id</th>
            <th>message</th>
          </tr>
          {props.fortunes.map((fortune) => (
            <tr>
              <td>{fortune.id}</td>
              <td>{fortune.message}</td>
            </tr>
          ))}
        </table>
      </body>
    </html>
  );
}
