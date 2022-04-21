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

const arr = [
  412, 214, 454, 46, 427, 50, 238, 219, 117, 211, 414, 111, 398, 424, 83, 142,
  185, 265, 313, 427, 498, 90, 460, 189, 88, 240, 468, 175, 355, 216, 469, 315,
  104, 74, 292, 74, 289, 120, 100, 399, 192, 234, 274, 70, 68, 251, 148, 291,
  72, 215, 35, 271, 61, 118, 41, 343, 408, 361, 222, 356, 96, 179, 113, 279,
  100, 0, 279, 17, 162, 246, 129, 376, 239, 83, 423, 237, 9, 493, 450, 310, 37,
  361, 296, 90, 71, 5, 303, 53, 369, 31, 56, 353, 309, 376, 69, 96, 415, 34, 66,
  265, 190, 49, 46, 301, 186, 53, 200, 396, 428, 137, 279, 462, 62, 64, 165,
  229, 466, 145, 95, 50, 351, 147, 195, 245, 22, 337, 80, 20, 338, 209, 446,
  478, 402, 410, 243, 97, 109, 411, 0, 236, 476, 341, 186, 8, 278, 143, 302,
  238, 315, 141, 123, 214, 36, 322, 494, 113, 480, 16, 359, 32, 113, 349, 16,
  20, 83, 152, 384, 196, 373, 256, 262, 407, 495, 157, 235, 220, 389, 172, 51,
  471, 225, 0, 224, 185, 240, 108, 30, 5, 256, 34, 70, 414, 492, 126, 82, 412,
  102, 348, 305, 165, 392, 214, 148, 474, 484, 4, 338, 235, 60, 277, 409, 264,
  330, 224, 190, 178, 32, 268, 186, 1, 181, 156, 339, 278, 428, 405, 102, 218,
  213, 479, 199, 249, 338, 156, 423, 473, 405, 199, 116, 261, 173, 66, 205, 389,
  301, 132, 474, 184, 119, 381, 213, 35, 460, 407, 371, 483, 299, 212, 168, 94,
  46, 208, 205, 301, 235, 47, 286, 179, 436, 237, 115, 102, 497, 230, 238, 81,
  443, 338, 364, 361, 284, 89, 105, 414, 29, 280, 458, 152, 434, 353, 143, 97,
  92, 32, 374, 158, 194, 170, 300, 375, 436, 413, 423, 434, 416, 250, 474, 284,
  369, 21, 423, 415, 97, 253, 373, 207, 443, 59, 296, 89, 471, 12, 38, 439, 353,
  56, 25, 419, 1, 322, 92, 132, 347, 135, 380, 211, 156, 358, 408, 494, 425,
  132, 173, 258, 480, 422, 84, 15, 385, 451, 86, 393, 472, 269, 366, 309, 102,
  440, 359, 186, 344, 172, 430, 133, 170, 457, 196, 321, 494, 409, 294, 476, 95,
  435, 49, 373, 413, 257, 332, 262, 65, 233, 386, 181, 0, 16, 379, 207, 445,
  487, 63, 463, 209, 41, 57, 440, 21, 193, 30, 42, 116, 481, 375, 388, 205, 208,
  281, 299, 93, 140, 305, 146, 296, 481, 271, 319, 11, 78, 62, 183, 310, 388,
  312, 403, 97, 63, 398, 40, 242, 307, 139, 403, 121, 425, 362, 442, 124, 2,
  285, 198, 195, 199, 200, 467, 46, 73, 289, 478, 452, 42, 285, 13, 472, 98, 38,
  116, 62, 48, 206, 249, 282, 309, 113, 443, 24, 379, 66, 455, 214, 347, 47,
  335, 31, 28, 217, 438, 432, 261, 496, 338, 161, 180, 350, 220, 306, 62, 418,
  185, 21, 429, 49, 317, 439, 380, 472, 308, 229, 301, 22, 371, 46, 410, 441,
  221, 78, 53, 209, 127, 12, 435, 446, 257, 140, 208, 47, 308, 391, 436, 189,
  23, 306, 286, 216, 50, 95, 322, 68, 400, 194, 236, 444, 99, 113, 299, 196,
  340, 356, 64, 266, 95, 144, 332, 487, 256, 367, 29, 299, 345, 333, 20, 167,
  458, 336, 2, 456, 48, 44, 382, 409, 274, 243, 422, 439, 146, 18, 145, 453,
  298, 402, 348, 153, 416, 497, 374, 398, 182, 231, 410, 256, 254, 53, 139, 241,
  340, 331, 316, 182, 140, 210, 338, 160, 192, 84, 487, 0, 337, 136, 403, 54,
  10, 113, 424, 336, 96, 224, 213, 380, 75, 336, 141, 99, 252, 90, 198, 259,
  348, 194, 399, 92, 473, 417, 139, 440, 14, 161, 189, 352, 414, 268, 295, 66,
  475, 40, 338, 147, 146, 312, 498, 328, 94, 298, 82, 278, 213, 279, 320, 50,
  289, 148, 331, 110, 321, 87, 492, 283, 315, 2, 394, 417, 226, 422, 424, 147,
  346, 106, 442, 56, 281, 125, 294, 403, 230, 477, 382, 88, 268, 78, 481, 434,
  21, 397, 111, 273, 313, 160, 159, 136, 34, 286, 96, 100, 341, 451, 310, 105,
  45, 479, 191, 244, 94, 60, 251, 434, 222, 52, 392, 17, 408, 179, 338, 236,
  289, 414, 47, 392, 220, 205, 36, 403, 294, 387, 348, 43, 220, 28, 172, 58, 67,
  154, 281, 10, 154, 346, 342, 220, 399, 469, 172, 349, 119, 368, 468, 324, 314,
  380, 264, 161, 258, 102, 14, 180, 96, 416, 459, 217, 472, 40, 297, 173, 170,
  23, 449, 292, 4, 395, 394, 357, 53, 390, 340, 226, 318, 234, 454, 483, 472,
  475, 285, 95, 117, 235, 429, 81, 327, 329, 46, 219, 81, 33, 354, 79, 386, 261,
  80, 465, 447, 453, 228, 140, 219, 307, 171, 340, 72, 71, 430, 19, 372, 437,
  265, 496, 143, 72, 8, 419, 271, 106, 102, 381, 118, 146, 141, 315, 147, 152,
  211, 221, 152, 498, 257, 389, 160, 190, 64, 489, 174, 168, 282, 139, 116, 297,
  338, 256, 222, 488, 99, 192, 373, 270, 495, 63, 476, 149, 426, 337, 279, 27,
  216, 335, 396, 259, 139, 488, 369, 173, 285, 49, 97, 417, 146, 108, 474, 159,
  236, 455, 176, 191, 110, 165, 26, 158, 254, 186, 124, 494, 218, 25, 106, 407,
  45, 273, 87, 63, 456, 451, 15, 340, 207, 387, 287, 97, 384, 252, 301, 354,
  264, 149, 106, 253, 289, 331, 479, 81, 197, 62, 94, 296, 319, 288, 84, 282,
  283, 357, 192, 478, 37, 434, 50, 79, 369, 437, 194, 297, 225, 177, 17, 117,
  394, 480, 135, 393, 168, 20, 308, 86, 219, 390, 397, 381, 286, 400, 464, 234,
  253, 245, 84, 333, 475, 403, 377, 373, 323, 49, 447, 399, 406, 466, 63, 426,
  186, 292, 324, 285, 150, 180, 194, 335, 271, 96, 185, 45, 417, 182, 431, 427,
  352, 479, 96, 134, 52, 57, 460, 88, 118, 451, 28, 300, 317, 152, 401,
];

console.time("bubbleSort");
const sorted = bubbleSort(arr);
console.timeEnd("bubbleSort");

console.log(sorted);
