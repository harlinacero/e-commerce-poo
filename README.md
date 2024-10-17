
## Cambios Realizados desde el Último Commit

### 1. Reorganización de la Estructura del Proyecto

- **Producto**: La clase `Product` se ha movido a una nueva carpeta llamada `producto`.
- **Nuevas Clases Derivadas**: Se han creado dos nuevas clases derivadas de `Product`:
  - `DigitalProduct`: Incluye atributos adicionales como `fileFormat` y `size`.
  - `PhysicalProduct`: Incluye atributos adicionales como `weight` y `dimensions`.

### 2. Implementación de la Interfaz `ProductInfo`

- **Interfaz `ProductInfo`**: Se ha definido una interfaz `ProductInfo` que declara el método `ShowInfo`.
- **Implementación en Clases Derivadas**: Las clases `DigitalProduct` y `PhysicalProduct` implementan la interfaz `ProductInfo` y su método `ShowInfo`.

### 3. Actualización de `ShoppingCar`

- **Uso de la Interfaz `ProductInfo`**: La clase `ShoppingCar` ahora utiliza la interfaz `ProductInfo` para manejar los productos y mostrar su información.
- **Método `ShowShoppingCar`**: Se ha actualizado el método `ShowShoppingCar` para utilizar el método `ShowInfo` de la interfaz `ProductInfo`.

### 4. Actualización del Archivo `main.go`

- **Importaciones**: Se han actualizado las importaciones para reflejar la nueva estructura del proyecto.
- **Interacción con el Usuario**: Se ha añadido lógica para interactuar con el usuario, permitiendo agregar productos al carrito de compras y mostrar la información del carrito.

### 5. Actualización de la Clase `User` y sus Derivadas

- **Clase `User`**: Se ha definido una interfaz `User` que declara los métodos `GetUsername` y `GetRole`.
- **Clase `BaseUser`**: Implementa los métodos comunes de la interfaz `User`.
- **Nuevas Clases Derivadas**: Se han creado tres nuevas clases derivadas de `BaseUser`:
  - `Admin`: Representa a un usuario administrador.
  - `Seller`: Representa a un usuario vendedor.
  - `Customer`: Representa a un usuario cliente.



- **domain/**: Contiene las clases de dominio del proyecto e-commerce.
- **product/**: Contiene las definiciones de los productos.
    - `digitalProduct.go`: Define la estructura y métodos relacionados con los productos digitales.
    - `physicalProduct.go`: Define la estructura y métodos relacionados con los productos físicos.
    - `product.go`: Define la estructura y métodos comunes de los productos.
    - `productinfo.go`: Define la interfaz `ProductInfo`.


- **shoppingcar/**: Contiene la definición del carrito de compras.
    - `shoppingCar.go`: Define la estructura y métodos relacionados con el carrito de compras.
    - `user.go`: Define la estructura y métodos relacionados con los usuarios.


   - `order.go`: Define la estructura y métodos relacionados con la generación de un pedido. Sus atributos son usuario, lista de productos, estado del pedido, fecha y total de la compra. Sus métodos son UpdateState, GetState, CancelOrder, GetOrder y GetTotal, para gestionar los pedidos.


   - `pay.go`: Define la estructura y métodos relacionados con el pago. Sus atributos son: id, método de pago, el pedido y el valor total del pago (es un atributo adicional al de la orden por que, dependiendo del método de pago pueden crearse tarifas adicionales o descuentos). Solo contiene el método procesar pago.


   - `review.go`: Define la estructura y métodos relacionados con los comentarios y calificaciones del producto. Sus atributos son usuario,  producto, comentario y puntuación. Sus métodos son: Agregar comentario, Agregar calificación.

**user/**: Contiene las definiciones de los usuarios.
  - `userbase.go`: Define la estructura y métodos relacionados para los usuarios comunes con los usuarios, por ahora solo setter y getter para poder acceder a sus atributos.
  - `admin.go`: Define la estructura y métodos relacionados con los usuarios de tipo administrador
  - `seller.go`: Define la estructura y métodos relacionados con los usuarios de tipo vendedor
  - `customer.go`: Define la estructura y métodos relacionados con los usuerios de tipo cliente
   

- **application/**: Contiene los servicios de la aplicación.
  - **services/**: Contiene los servicios para agregar usuarios y productos.
    - `addUser.go`: Servicio para agregar un nuevo usuario.
    - `addNewProduct.go`: Servicio para agregar un nuevo producto.

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

