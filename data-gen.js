const getRandomInt = (min, max) => {
  return Math.floor(Math.random() * (max - min + 1)) + min;
};

const numberOfNumbers = 100_000;
const minNumber = 1;
const maxNumber = 1_000_000;
const numbers = [];

for (let i = 0; i < numberOfNumbers; i++) {
  const randomNumber = getRandomInt(minNumber, maxNumber);
  numbers.push(randomNumber);
}

const jsonData = JSON.stringify(numbers);
const fileName = `data.json`;

Bun.write(fileName, jsonData);
