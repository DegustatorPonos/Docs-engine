// Determines the parsing mode whe engine runs on
let Mode = 0;


async function DisplayContent(path) {
	let returnStrings = [];
	await GetString("ReadFile?path=" + path	).then(result => {
		returnStrings = result.split('\n');
	}); 
	DetermineParsingMode(returnStrings[0])
	console.log("Parsing mode: " + Mode);
	returnStrings.forEach(str => {
		if(str == null || str == "") return;
		console.log(str);
	});
}

// First returning string contains the info regarding parsing
function DetermineParsingMode(FirstString) {
	let _mode = FirstString.replace("mode ", "");
	Mode = _mode;
}
