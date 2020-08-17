

// function loadNav() {
//     var includeTals = document.getElementsByTagName("include");
    

// }

$(function(){
    $("#include-nav").load("nav.html");
  });
  
// replaceIncludeElements: function() {
//     var $this = this;
//     var filePath = $this.getFilePath();
//     var includeTals = document.getElementsByTagName("include");
//     this.forEach(includeTals, function() {
//         //拿到路径
//         var src = this.getAttribute("src");
//         //拿到文件内容
//         var content = $this.getFileContent($this.getRequestUrl(filePath, src));
//         //将文本转换成节点
//         var parent = this.parentNode;
//         var includeNodes = $this.parseNode($this.getHtml(content));
//         var size = includeNodes.length;
//         for(var i = 0; i < size; i++) {
//             parent.insertBefore(includeNodes[0], this);
//         }
//         //执行文本中的额javascript
//         $this.executeScript(content);
//         parent.removeChild(this);
//         //替换元素 this.parentNode.replaceChild(includeNodes[1], this);
//     })
// }
