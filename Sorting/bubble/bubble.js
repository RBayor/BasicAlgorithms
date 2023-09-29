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

const t0 = performance.now();
const arr = [];
for (let i = 0; i < 10_000; i++) {
  arr.push(Math.floor(Math.random() * 100000));
}

const sorted = bubbleSort(arr);
const t1 = performance.now();

const time = ((t1 - t0) / 1000).toFixed(4);
console.table({
  "Sort Type": "Bubble Sort",
  "Time elapsed": `${time} seconds`,
  "Items length": arr.length.toLocaleString(),
});
