import { readFile } from "node:fs/promises";

async function part2(path) {
  const file = await readFile(path, "utf8");
  const data = file.trim().split("\n");

  const adj = createAdj(data);
  const memo = new Map();

  function dfs(s, f, d) {
    const key = `${s}|${f}|${d}`;
    if (memo.has(key)) return memo.get(key);

    if (s === "out") {
      const res = f && d ? 1 : 0;
      memo.set(key, res);
      return res;
    }

    let total = 0;

    adj[s].forEach((el) => {
      const nf = f || el === "fft";
      const nd = d || el === "dac";
      total += dfs(el, nf, nd);
    });

    memo.set(key, total);
    return total;
  }

  const sum = dfs("svr", false, false);
  console.log("Day 11 | Part 2:", sum);
}

function createAdj(data) {
  const adj = {};

  data.forEach((el) => {
    const parts = el.split(":");
    adj[parts[0]] = parts[1].trim().split(" ");
  });

  return adj;
}

part2("input.txt", false, false);
