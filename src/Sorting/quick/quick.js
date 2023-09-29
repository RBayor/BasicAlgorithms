/**
 * Implementation of Quick sort algorithm.
 * @param {Array} arr
 * @returns sorted array (asc)
 */
const quickSort = (arr) => {
  if (arr.length <= 1) return arr;

  const pivot = arr[0];
  const left = [];
  const right = [];
  for (let i = 1; i < arr.length; i++) {
    if (arr[i] < pivot) {
      left.push(arr[i]);
    } else {
      right.push(arr[i]);
    }
  }
  return [...quickSort(left), pivot, ...quickSort(right)];
};

const arr = [];
for (let i = 0; i < 50_000; i++) {
  arr.push(Math.floor(Math.random() * 100000));
}

const t0 = performance.now();
const sorted = quickSort(arr);
const t1 = performance.now();

const sortTime = ((t1 - t0) / 1000).toFixed(4);

console.table({
  "Sort Type": "Quick Sort",
  "Time elapsed": `${sortTime} seconds`,
  "Items length": arr.length.toLocaleString(),
});
