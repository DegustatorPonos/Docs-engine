const directoriesDictionanry = {};

let mutex = 0;

async function UpdateStructure() {
    panel = document.getElementById("NavPanel");
	 //Goofy ahh synchronization method
	await GetDirectoriesArray('/')
	let mutexPromise = new Promise((resolve) => { 
			setInterval(() => {
				if(mutex == 0) {
					resolve(""); 
					return;
				}
			}, 1);
	});
	await mutexPromise.then(() => { 
		console.log(directoriesDictionanry);
		console.log(directoriesDictionanry["/Decoration/"]);
		DrawTree('/');
		});
}


//Fills up global dictionary variable recursively
async function GetDirectoriesArray(rootDir) {
	let currentDirs = '';
	mutex += 1;
	console.log("Mutex: " + mutex);
	await GetString("GetDirectories?path=" + rootDir).then(result => {
		currentDirs = result;
	});
	directoriesDictionanry[rootDir] = currentDirs;
	currentDirs.split(';').forEach(async el => {
		//Recursion
		if(el.includes(' -d')) {
			await GetDirectoriesArray(rootDir+el.replace(' -d', '')+'/');
		}
	});
	mutex -= 1;
	console.log("Mutex -" + mutex)
}


//Sync recursive function to draw a dir tree
function DrawTree(baseDir) {
	console.log("Keys = " + Object.keys(directoriesDictionanry)[1]);
	console.log('Requested listing "' + baseDir + '"');
	let contents = directoriesDictionanry[baseDir];
	console.log('Got from dict ' + contents);
	contents.split(';').forEach(el => {
		if(el.includes(' -f')) {
			AddFile(new File(el.replace(' -f', ''), baseDir));
		}

		if(el.includes(' -d')) {
			AddDir(new Directory(el.replace(' -d', ''), baseDir));
			DrawTree(baseDir + el.replace(' -d', '') + '/');
		}
	});
	console.log('Content > ' + contents);	
}


//=========Page update==========

//Adds file to the view
function AddFile(file) {
    let displayElement = document.createElement("p");
    displayElement.innerHTML = "| " + file.name ;
    displayElement.className = "FileElement FileElementOverride";
    //We redirect 
    displayElement.id = file.path + "/" + file.name + ".md";
    let nav = document.getElementById('NavPanel');
    nav.appendChild(displayElement);
}

//Adds directory to the view. Can be recursive.
function AddDir(dir) {
    if(dir == undefined) return;
    let displayElement = document.createElement("p");
    displayElement.innerHTML = dir.name + " dir";
    let nav = document.getElementById('NavPanel');
    nav.appendChild(displayElement);
}


//==========CLASSES============
class File {
    displayName = "";
    name = "";
    path = "";
    constructor(name, path) {
        this.name = name;
        this.path = path;
    }
}

class Directory {
    name = "";
    path = "";
    constructor(name, path) {
        this.name = name;
        this.path = path;
    }
}
