/**
 * Implementation of a hash table
 */
class HashTable {
  constructor() {
    this.table = new Array(17011);
    this.numOfItems = 0;
    this.load = 0;
  }

  #hashKeyToIndex = (key) => {
    let hash = 719;

    for (let i = 0; i < key.length; i++) {
      hash = (83 * hash * key.charCodeAt(i)) % this.table.length;
    }

    return hash;
  };

  #resize = () => {
    const currenTable = this.table;
    const newTable = Array(this.#nextPrime(this.table.length * 2));

    console.log(
      `Resized from '${currenTable.length}' to '${newTable.length}'\n`
    );

    this.numOfItems = 0;
    this.table = newTable;

    currenTable.forEach((item) => {
      if (item) {
        item.forEach(([key, value]) => {
          this.#storeValue(key, value);
        });
      }
    });
  };

  setItem = (key, value) => {
    this.load = this.numOfItems / this.table.length;
    const index = this.#hashKeyToIndex(key);

    if (this.load > 0.8 && this.table[index]) {
      this.#resize(key, value);
    }

    this.#storeValue(key, value);

    return;
  };

  #storeValue = (key, value) => {
    const index = this.#hashKeyToIndex(key);

    if (this.table[index]) {
      const item = this.table[index].find((x) => x[0] === key);

      if (item) {
        item[1] = value;
        return;
      } else {
        this.table[index].push([key, value]);

        this.numOfItems++;
        return;
      }
    }

    this.numOfItems++;
    this.table[index] = [[key, value]];
  };

  getItem = (key) => {
    const index = this.#hashKeyToIndex(key);

    if (!this.table[index]) return null;

    return this.table[index].find((x) => x[0] === key)[1];
  };

  #isPrime = (num) => {
    let sqrtnum = Math.floor(Math.sqrt(num));
    let prime = num !== 1;
    for (let i = 2; i < sqrtnum + 1; i++) {
      if (num % i === 0) {
        prime = false;
        break;
      }
    }
    return prime;
  };

  #nextPrime = (num) => {
    if (!num) return;
    while (!this.#isPrime(++num)) {}
    return num;
  };
}

const t = new HashTable();

console.time("hashTable");
console.time("setItem");
// t.setItem("homer-1", "first");
// t.setItem("bart", "second");
// t.setItem("lisa", "third");
// t.setItem("marge", "fifth");
// t.setItem("homer", "fourth");
// t.setItem("krusty", "sixth");
// t.setItem("burns", "seventh");
// t.setItem("burns", "choco");
for (let i = 0; i <= 1000000; i++) {
  t.setItem(`item-${i}`, `data-${i}`);
}
console.timeEnd("setItem");

// console.time("getItem");
// t.getItem("homer");
// t.getItem("bart");
// t.getItem("bart");
// t.getItem("lisa");
// t.getItem("marge");
// t.getItem("krusty");
// t.getItem("burns");
// t.getItem("burns");
// console.timeEnd("getItem");
console.timeEnd("hashTable");

// console.table(t.table);
console.table({
  "table size": t.table.length,
  "num of items:": t.numOfItems,
  load: t.load,
});
