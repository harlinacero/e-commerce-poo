

- **domain/**: Contiene las clases de dominio del proyecto e-commerce.
  - `product.go`: Define la estructura y métodos relacionados con los productos Sus atributos son: id, name, description, grossvalue, taxaddpercentage, discpercentage, y stock, además, de los métodos setter y getter para poder acceder a sus atributos, cuenta con los métodos GetDiscValue, GetTaxAddValue y GetRealValue donde se realiza la lógica correspondiente para calcular estos valores.
  - `shoppingCar.go`: Define la estructura y métodos relacionados con el carrito de compras, tiene como atributos el usuario y una lista de productos. Sus métodos son AddProduct, RemoveProduct, GetProducts y CalculateTotalValue para administrar la lista de productos del carrito.
   - `order.go`: Define la estructura y métodos relacionados con la generación de un pedido. Sus atributos son usuario, lista de productos, estado del pedido, fecha y total de la compra. Sus métodos son UpdateState, GetState, CancelOrder, GetOrder y GetTotal, para gestionar los pedidos.
   - `pay.go`: Define la estructura y métodos relacionados con el pago. Sus atributos son: id, método de pago, el pedido y el valor total del pago (es un atributo adicional al de la orden por que, dependiendo del método de pago pueden crearse tarifas adicionales o descuentos). Solo contiene el método procesar pago.
   - `review.go`: Define la estructura y métodos relacionados con los comentarios y calificaciones del producto. Sus atributos son usuario,  producto, comentario y puntuación. Sus métodos son: Agregar comentario, Agregar calificación.
  - `user.go`: Define la estructura y métodos relacionados con los usuarios, por ahora solo setter y getter para poder acceder a sus atributos.
- **go.mod**: Archivo de módulos de Go que define las dependencias del proyecto.
- **main.go**: Archivo principal que contiene la lógica para interactuar con el usuario y gestionar productos.

## Ejecución del Proyecto

Para ejecutar el proyecto, sigue estos pasos:

1. **Clonar el repositorio**:
    ```sh
    git clone <URL_DEL_REPOSITORIO>
    cd <NOMBRE_DEL_REPOSITORIO>
    ```

2. **Instalar dependencias**:
    ```sh
    go mod tidy
    ```

3. **Ejecutar la aplicación**:
    ```sh
    go run main.go
    ```

## Uso

Al ejecutar la aplicación, se te pedirá que introduzcas la información del producto:

1. Nombre del producto
2. Descripción del producto
3. Valor bruto del producto
4. Porcentaje de impuesto
5. Porcentaje de descuento
6. Cantidad de stock

La aplicación calculará y mostrará los valores correspondientes al impuesto, descuento y valor real del producto.
![image](https://github.com/user-attachments/assets/b29bdb03-002b-4b4c-92da-69253e3ecd37)

