class HashTable {
    constructor(size = 7) {
        this.dataMap = new Array(size);
    }
   
    _hash(key) {
        let hash = 0;
        for (let i = 0; i < key.length; i++) {
            hash = (hash + key.charCodeAt(i) * 23) % this.dataMap.length;
        }
        return hash;
    }

    printTable() {
        for (let i = 0; i < this.dataMap.length; i++) {
            console.log(i, ": ", this.dataMap[i]);
        }
    }

    set(key,value) {
        let index = this._hash(key);
        if(!this.dataMap[index]) this.dataMap[index] = [];
        
        this.dataMap[index].push([key, value]);
        return this;
    }

	get(key) {
        const index = this._hash(key);
        const node = this.dataMap[index];

        if (node === undefined) {
            return undefined;
        }

        for (let i = 0; i < node.length; i++) {
            if (node[i][0] === key) {
                return node[i][1];
            }
        }

        return undefined;
    }

}



let myHashTable = new HashTable();

myHashTable.set("nails", 100);
myHashTable.set("tile", 50);
myHashTable.set("lumber", 80);

console.log("Lumber:");
console.log( myHashTable.get("lumber") );

console.log("\nBolts:");
console.log( myHashTable.get("bolts") );


/*
    EXPECTED OUTPUT:
    ----------------
    Lumber:
    80

    Bolts:
    undefined

*/