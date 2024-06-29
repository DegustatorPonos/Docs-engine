async function UpdateStructure() {
    let resp = await GetString("GetDirectories"); 

    let parsedResponce = await ParseDirectory('', resp)

    console.log(parsedResponce);  
}

//Recursive 
async function ParseDirectory(rootPath, content) {
    let returnArray = content.split(';');
    let outputaArray = [];
    returnArray.forEach(async element => {
        if(element.includes(' -f')) {
            element = element.replace(" -f", "");
            outputaArray.push(new File(element, rootPath));
        }
        else if(element.includes(' -d')) {
            element = element.replace(" -d", "");
            let dirPath = rootPath + "/" + element;
            let dirContent = await GetString("GetDirectories?path=" + dirPath);
            console.log("Requested path " + dirPath)
            outputaArray.push(await ParseDirectory(dirPath, dirContent));
        }
    });
    return outputaArray;
}

//==========CLASSES============
class File {
    name = "";
    path = "";
    constructor(name, path) {
        this.name = name;
        this.path = path;
    }
}