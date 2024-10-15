

- **domain/**: Contiene los archivos de dominio del proyecto.
  - `product.go`: Define la estructura y métodos relacionados con los productos.
  - `shoppingCar.go`: Define la estructura y métodos relacionados con el carrito de compras.
  - `user.go`: Define la estructura y métodos relacionados con los usuarios.
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

