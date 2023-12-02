// Note: reuse_port is available only in Linux.
// See: https://github.com/oven-sh/bun/issues/2428#issuecomment-1750544989

import os from 'node:os';

const numCPUs = os.cpus().length;
for (let i = 0; i < numCPUs; i++) {
  Bun.spawn(['bun', 'index.ts'], {
    stdio: ['inherit', 'inherit', 'inherit'],
    env: { ...process.env },
  });
}
