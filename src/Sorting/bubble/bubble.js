const fs = require("fs").promises;
const path = require("path");

const runBenchmark = async () => {
  const numIterations = 1;
  const creationTimes = [];
  const sortTimes = [];
  const messyArr = await readSortableArray();

  if (messyArr.length === 0) {
    console.log("Failed to read file");
    return;
  }

  for (let i = 0; i < numIterations; i++) {
    const { creationTime, sortTime } = benchmarkBubbleSort([...messyArr]);
    creationTimes.push(creationTime);
    sortTimes.push(sortTime);
  }

  const AVG_CREATION_TIME =
    creationTimes.reduce((a, b) => +a + +b) / creationTimes.length;
  const AVG_SORT_TIME = sortTimes.reduce((a, b) => +a + +b) / sortTimes.length;

  const table = {
    "Sort Type": "Bubble Sort",
    "Array Size": messyArr.length.toLocaleString(),
    Iterations: numIterations.toLocaleString(),
    ____________a____________: "____________a____________",
    "Avg Creation Time (ms)": `${AVG_CREATION_TIME.toFixed(4)}ms`,
    "Avg Sort Time (ms)": `${AVG_SORT_TIME.toFixed(4)}ms`,
    ____________b____________: "____________b____________",
    "Avg Creation Time (s)": ` ${(AVG_CREATION_TIME / 1000).toFixed(4)}s`,
    "Avg Sort Time (s)": `${(AVG_SORT_TIME / 1000).toFixed(4)}s`,
    ____________c____________: "____________c____________",
    "Creation Times (s)": creationTimes
      .map((t) => (t / 1000).toFixed(4))
      .toLocaleString(),
    "Sort Times (s)": sortTimes
      .map((t) => (t / 1000).toFixed(4))
      .toLocaleString(),
  };

  console.table(table);
};

const measureTime = (operation) => {
  const t0 = performance.now();
  operation();
  const t1 = performance.now();
  return (t1 - t0).toFixed(4);
};

const benchmarkBubbleSort = (messyArr) => {
  const size = messyArr.length;
  const creationTime = measureTime(() => generateRandomArray(size));
  const sortTime = measureTime(() => bubbleSort(messyArr));

  return { creationTime, sortTime };
};

const generateRandomArray = (size) => {
  const arr = [];
  for (let i = 0; i < size; i++) {
    arr.push(Math.floor(Math.random() * 1000000));
  }
  return arr;
};

const readSortableArray = async () => {
  const filePath = path.join(__dirname, "../../../data.json");

  try {
    const res = JSON.parse(await fs.readFile(filePath, "utf8"));

    return structuredClone(res);
  } catch (err) {
    console.log(err);
  }

  return [];
};

/**
 * Implementation of optimized Bubble Sort algorithm
 * @param {Array} array
 * @returns sorted array (asc)
 */
const bubbleSort = (arr) => {
  let swapped;
  do {
    swapped = false;
    for (let i = 0; i < arr.length; i++) {
      if (arr[i] > arr[i + 1]) {
        [arr[i], arr[i + 1]] = [arr[i + 1], arr[i]];
        swapped = true;
      }
    }
  } while (swapped);
  return arr;
};

runBenchmark();
