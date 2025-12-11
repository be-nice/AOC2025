import { readFile } from "node:fs/promises";

async function part1(path) {
  const file = await readFile(path, "utf8");
  const data = file.trim().split("\n");

  const adj = createAdj(data);

  const dfs = (s) => {
    if (s === "out") {
      return 1;
    }

    let total = 0;

    adj[s].forEach((el) => {
      total += dfs(el);
    });

    return total;
  };

  console.log(dfs("you"));
}

function createAdj(data) {
  const adj = {};

  data.forEach((el) => {
    const parts = el.split(":");
    adj[parts[0]] = parts[1].trim().split(" ");
  });

  return adj;
}

part1("input.txt");
