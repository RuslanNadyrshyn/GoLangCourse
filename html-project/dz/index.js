
function clearProducts() {
    let list = document.querySelector('.list');
    if (list != null) {
        while (list.hasChildNodes()) {
            list.removeChild(list.firstChild);
        }
    }
}

function displayProducts(array, start) {
    let list = document.querySelector('.list');
    for (let i = start; i < array.length; i++) {
        let listItem = document.createElement('div');
        listItem.className = 'list_item';
        list.appendChild(listItem);

        let productListImg = document.createElement('div');
        productListImg.className = 'list_img';
        productListImg.innerHTML = `<img class="list_logo" src="${array[i].image}">`;
        listItem.appendChild(productListImg);

        let productListText = document.createElement('div');
        productListText.className = 'product_list_text';
        productListText.innerHTML = array[i].name;
        listItem.appendChild(productListText);

        let productListPrice = document.createElement('div');
        productListPrice.className = 'product_list_price';
        productListPrice.innerHTML = `${array[i].price} $`;
        listItem.appendChild(productListPrice);
    }
}

const loadElements = () => new Promise((res) =>{
    fetch("http://localhost:8080/get_products", {
        method: "GET"
    }).then(data => data.json().then(json_data => res(json_data)))
})

const addProduct = async () => {
    let newMenuId = document.getElementById('product_menu_id').value;
    let newName= document.getElementById('product_name').value;
    let newPrice = document.getElementById('product_price').value;
    let newImage = document.getElementById('product_image').value;
    let newType = document.getElementById('product_type').value;
    let newIngredients = document.getElementById('product_ingredients').value;
    newIngredients = JSON.stringify(newIngredients);
    let ingredients = newIngredients.split(' ');
    console.log(newMenuId, newName, newPrice, newImage, newType, newIngredients);
    try {
        const response = await fetch("http://localhost:8080/add_product", {
            method: "POST",
            body: JSON.stringify({
                MenuId: newMenuId,
                Name: newName,
                Price: newPrice,
                Image: newImage,
                Type: newType,
                Ingredients: ingredients
            }),
            headers: {
                'Content-type': 'application/json; charset=UTF-8',
            },
        })
        const json = await response.json();
        console.log('Успех:', JSON.stringify(json));
        return json;
    } catch (error) {
        console.error('Ошибка:', error);
    }
}

const main = async () => {
    let array = await loadElements();
    clearProducts();                 // Очищаем вывод
    displayProducts(array, 0); // Отображаем всё заново

    localStorage.setItem("array", JSON.stringify(array));
}

const loop = () => {
    new Promise((resolve) => {
        main();
        setTimeout(() => {
            resolve()
        }, 1000)
     }).then(loop)
}
 loop()