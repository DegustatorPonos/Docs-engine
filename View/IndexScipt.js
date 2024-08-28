//Relace it with your logic server's address
const baseAddress = "http://127.0.0.1:5000";

async function Init() {
    //I just don't like how JQuerry looks like so I separate them

    //Update the right directory menu
    await UpdateStructure();
}



async function GetString(uri) {
    let responce = await fetch(baseAddress + "/" + uri);
    let output = await responce.text().then(outp => {
	return outp;
    });
    //console.log(output); LOGGING 
    return output;
}

//Occurs when the element on the side panel is selected
async function GoToPage(node) {
    console.log(node.id);
}


