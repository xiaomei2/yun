let image = '';

function selectImage(file){
    if(!file.files || !file.files[0]){
        return;
    }
    var reader = new FileReader();
    reader.onload = function(evt){
        document.getElementById('img_preview').src = evt.target.result;
        image = evt.target.result;
        console.log(image)
    }
    reader.readAsDataURL(file.files[0]);
}
function uploadImage() {
    if (image === '') {
        alert('请先选择图片');
        return;
    }

    // 将图片数据转换成 Base64 格式
    const base64Image = image.split(',')[1]; // 去除 data:image/jpeg;base64, 这样的前缀部分

// 构建包含图片 Base64 数据的 JSON 对象
    const imageData = {
        imagesCode: base64Image
    };

// 发送 JSON 数据到服务器，使用 Fetch API 发送 POST 请求
    fetch('http:/localhost:8080/pickupInfo', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(imageData)
    })
        .then(response => response.json())
        .then(data => {
            console.log('上传成功:', data);
        })
        .catch(error => {
            console.error('上传失败:', error);
        });

}