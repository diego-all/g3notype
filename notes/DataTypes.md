# Golang DataTypes


Go tiene los siguientes tipos de enteros independientes de la arquitectura:

    uint8       unsigned  8-bit integers (0 to 255)
    uint16      unsigned 16-bit integers (0 to 65535)
    uint32      unsigned 32-bit integers (0 to 4294967295)
    uint64      unsigned 64-bit integers (0 to 18446744073709551615)
    int8        signed  8-bit integers (-128 to 127)
    int16       signed 16-bit integers (-32768 to 32767)
    int32       signed 32-bit integers (-2147483648 to 2147483647)
    int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)


Los números flotantes y complejos también vienen en diferentes tamaños:

    float32     IEEE-754 32-bit floating-point numbers
    float64     IEEE-754 64-bit floating-point numbers
    complex64   complex numbers with float32 real and imaginary parts
    complex128  complex numbers with float64 real and imaginary parts


También existen varios alias de tipos de números, que asignan nombres útiles a tipos de datos específicos:

    byte        alias for uint8
    rune        alias for int32


Además, Go tiene los siguientes tipos específicos de la implementación:

    uint     unsigned, either 32 or 64 bits
    int      signed, either 32 or 64 bits
    uintptr  unsigned integer large enough to store the uninterpreted bits of a pointer value


Y muchos mas...

    https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go-es


Por otro lado


    string
    bool


 # Resumen de tipos de datos seleccionados

    string
    int
    float32
    float64
    bool

