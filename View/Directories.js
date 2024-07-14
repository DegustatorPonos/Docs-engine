 function UpdateStructure() {
    panel = document.getElementById("NavPanel");
    GetString("GetDirectories").then(resp => {
        let parsedResponce = ParseDirectory('', resp);
        console.log(parsedResponce)
    }); 
    
    
    //AddDir(parsedResponce, panel);
    ;  
}

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

//Recursive function to retuern a file directory structure
function ParseDirectory(rootPath, content) {
    let returnArray = content.split(';');
    let outputArray = [];
    returnArray.forEach(element => {
        if(element.includes(' -f')) {
            //File parsing
            element = element.replace(" -f", "");
            outputFile = new File(element, rootPath);
            outputArray.push(outputFile);
            AddFile(outputFile);
        }
        else if(element.includes(' -d')) {
            //Directory parsing
            element = element.replace(" -d", "");
            let dirPath = rootPath + "/" + element;

            //Somehow it makes it somehow synchronious
            GetString("GetDirectories?path=" + dirPath).then(result => {
                AddDir(new Directory(element, dirPath));
                outputArray.push(ParseDirectory(dirPath, result));
            }).then(i => { return; });
        }
    });
    
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