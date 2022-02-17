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

const main = async () => {
    let array = await fetch("http://localhost:8080/get_products", { method: "GET"
    })
    array = await array.json();
    
    displayProducts(array, 0); // Отображаем всё заново
}

main();
