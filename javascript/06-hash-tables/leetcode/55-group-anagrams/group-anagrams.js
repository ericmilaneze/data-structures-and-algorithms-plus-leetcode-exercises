//   +=====================================================+
//   |                 WRITE YOUR CODE HERE                |
//   | Description:                                        |
//   | - This function groups anagrams from an array       |
//   |   of strings.                                       |
//   |                                                     |
//   | Return type: array                                  |
//   | - Returns an array of arrays where each array       |
//   |   contains anagrams.                                |
//   |                                                     |
//   | Tips:                                               |
//   | - You can use either a Map or an object to manage   |
//   |   the groups of anagrams.                           |
//   | - Example with Map:                                 |
//   |   anagramGroups.set(canonical, group);              |
//   | - Example with object:                              |
//   |   anagramGroups[canonical] = group;                 |
//   +=====================================================+


function groupAnagrams(anagrams) {
    const map = new Map();

    for (let i = 0; i < anagrams.length; i++) {
        const currentAnagram = anagrams[i];

        let keyFound = false;
        for (const key of map.keys()) {
            const innerMap = new Map();

            for (const char of key) {
                innerMap.set(char, (innerMap.get(char) || 0) + 1);
            }

            for (const char of currentAnagram) {
                innerMap.set(char, (innerMap.get(char) || 0) - 1);
            }

            let isPartOf = true;
            for (const v of innerMap.values()) {
                if (v !== 0) {
                    isPartOf = false;
                    break;
                }
            }

            if (isPartOf) {
                const list = map.get(key);
                list.push(currentAnagram);
                keyFound = true;
                break;
            }
        }
        
        if (!keyFound) {
            map.set(currentAnagram, [currentAnagram]);
        }
    }

    const arr = [];
    for (const key of map.keys()) {
        arr.push(map.get(key));
    }

    return arr;
}


// ---------------
// Lowercase Anagrams
// ---------------
console.log("Lowercase Anagrams:");
console.log("Input: ['eat', 'tea', 'tan', 'ate', 'nat', 'bat']");
console.log("Output: ", JSON.stringify(groupAnagrams(['eat', 'tea', 'tan', 'ate', 'nat', 'bat'])));
console.log("---------------");

// ---------------
// Mixed Case Anagrams
// ---------------
console.log("Mixed Case Anagrams:");
console.log("Input: ['Eat', 'Tea', 'Tan', 'Ate', 'Nat', 'Bat']");
console.log("Output: ", JSON.stringify(groupAnagrams(['Eat', 'Tea', 'Tan', 'Ate', 'Nat', 'Bat'])));
console.log("---------------");

// ---------------
// No Anagrams
// ---------------
console.log("No Anagrams:");
console.log("Input: ['hello', 'world', 'test']");
console.log("Output: ", JSON.stringify(groupAnagrams(['hello', 'world', 'test'])));
console.log("---------------");

// ---------------
// Empty Strings
// ---------------
console.log("Empty Strings:");
console.log("Input: ['', '', '']");
console.log("Output: ", JSON.stringify(groupAnagrams(['', '', ''])));
console.log("---------------");

// ---------------
// Single Characters
// ---------------
console.log("Single Characters:");
console.log("Input: ['a', 'b', 'a']");
console.log("Output: ", JSON.stringify(groupAnagrams(['a', 'b', 'a'])));
console.log("---------------");


// ---------------
// Special Characters
// ---------------
console.log("Single Characters:");
console.log("Input: ['123','321','!!','!!','$$']");
console.log("Output: ", JSON.stringify(groupAnagrams(['123','321','!!','!!','$$'])));
console.log("---------------");
