import { readFile } from "node:fs/promises";

async function part2(path) {
  const file = await readFile(path, "utf8");
  const data = file.trim().split("\n");

  const cols = data[0].length;
  let count = 0;

  let beams = Array(cols).fill(0);
  beams[data[0].indexOf("S")] = 1;

  let next = Array(cols).fill(0);

  for (const line of data.slice(1)) {
    next.fill(0);

    beams.forEach((n, col) => {
      if (n === 0) return;

      if (line[col] === "^") {
        col - 1 >= 0 ? (next[col - 1] += n) : (count += n);
        col + 1 < cols ? (next[col + 1] += n) : (count += n);
      } else {
        next[col] += n;
      }
    });

    [beams, next] = [next, beams];
  }

  count += beams.reduce((a, b) => a + b, 0);

  console.log(count);
}

part2("input.txt");
