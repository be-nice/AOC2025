import { readFile } from "node:fs/promises";

async function part1(path) {
  const file = await readFile(path, "utf8");
  const data = file.trim().split("\n");

  let count = 0;

  let beams = Array(data[0].length).fill(false);

  const start = data[0].indexOf("S");
  beams[start] = true;

  let next = Array(beams.length).fill(false);

  for (const line of data.slice(1)) {
    next.fill(false);

    beams.forEach((active, col) => {
      if (!active) return;

      if (line[col] === "^") {
        if (col - 1 >= 0) next[col - 1] = true;
        if (col + 1 < line.length) next[col + 1] = true;

        count++;
      } else {
        next[col] = true;
      }
    });

    [beams, next] = [next, beams];
  }

  console.log(count);
}

part1("input.txt");
