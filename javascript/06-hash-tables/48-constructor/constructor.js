class HashTable {
	constructor() {
        this.dataMap = new Array(7);
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

}



let myHashTable = new HashTable();
myHashTable.printTable();  

    
    /*
        EXPECTED OUTPUT:
        ----------------
        0 :  undefined
        1 :  undefined
        2 :  undefined
        3 :  undefined
        4 :  undefined
        5 :  undefined
        6 :  undefined

    */