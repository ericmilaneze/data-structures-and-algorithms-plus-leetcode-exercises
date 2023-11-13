class Node {
    constructor(value) {
        this.value = value;
        this.left = null;
        this.right = null;
    }
}
 
class BST {
    constructor() {
        this.root = null;
    }

	insert(value) {
        if (this.root === null) {
            this.root = new Node(value);
            return this;
        }

        let temp = this.root;

        while (temp !== null) {
            if (value === temp.value) {
                return undefined;
            } else if (value > temp.value) {
                if (temp.right === null) {
                    temp.right = new Node(value);
                    return this;
                } else {
                    temp = temp.right;
                }
            } else if (value < temp.value) {
                if (temp.left === null) {
                    temp.left = new Node(value);
                    return this;
                } else {
                    temp = temp.left;
                }
            }
        }

        return this;
    }

}



let myBST = new BST();

myBST.insert(2);
myBST.insert(1);
myBST.insert(3);
myBST.insert(3);

/*
    THE LINES ABOVE CREATE THIS TREE:
                 2
                / \
               1   3
*/


console.log("Root:", myBST.root.value);
console.log("\nRoot->Left:", myBST.root.left.value);
console.log("\nRoot->Right:", myBST.root.right.value);


/*
    EXPECTED OUTPUT:
    ----------------
    Root: 2

    Root->Left: 1

    Root->Right: 3

*/
      
