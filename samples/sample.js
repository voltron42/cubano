console.log(props.message)
console.log(file.cwd)
var fileprops = Object.keys(file)
console.log(fileprops)
fileprops.forEach(function(prop){
  console.log(typeof file[prop])
})
var hello = file.read("../hello.txt")
if (hello instanceof Error) {
  throw hello
}
console.log(hello)