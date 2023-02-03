package numtoletter

import (
	"errors"
	"strings"
)

var (
	//ErrValorNoAdmitido error para valor no admitidos
	ErrValorNoAdmitido = errors.New("Valor no admitido")
	us                 = []string{"cero", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"}
	ds                 = []string{"X", "y", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"}
	des                = []string{"diez", "once", "doce", "trece", "catorce", "quince", "dieciseis", "diecisiete", "dieciocho", "diecinueve"}
	cs                 = []string{"x", "cien", "doscientos", "trescientos", "cuatrocientos", "quinientos", "seiscientos", "setecientos", "ochocientos", "novecientos"}
)

// IntLetra convierte un numero entero a su representacion en palabras
func IntLetra(n int) (s string, err error) {
	var aux string
	sb := strings.Builder{}
	if n < 0 {
		sb.WriteString("menos")
		n *= -1
	}
	millones := quotient(n, 1000000)
	if millones > 999999 {
		return s, ErrValorNoAdmitido
	}
	millares := quotient(n, 1000) % 1000
	centenas := quotient(n, 100) % 10
	decenas := quotient(n, 10) % 10
	unidades := n % 10

	if millones == 1 {
		sb.WriteString(" un millón")
	} else if millones > 1 {
		aux, err = IntLetra(millones)
		sb.WriteString(aux)
		sb.WriteString(" millones")
	}
	if millares == 1 {
		sb.WriteString(" mil")
	} else if millares > 1 {
		aux, err = IntLetra(millares)
		sb.WriteRune(' ')
		sb.WriteString(aux)
		sb.WriteString(" mil")
	}
	if centenas == 1 {
		if n%100 == 0 {
			sb.WriteString(" cien")
		} else {
			sb.WriteString(" ciento")
		}
	} else if centenas > 0 {
		sb.WriteRune(' ')
		sb.WriteString(cs[centenas])
	}
	if decenas == 1 {
		sb.WriteRune(' ')
		sb.WriteString(des[n%10])
		unidades = 0
	} else if decenas == 2 && unidades > 0 {
		sb.WriteString(" veinti")
		sb.WriteString(us[unidades])
		unidades = 0
	} else if decenas > 1 {
		sb.WriteRune(' ')
		sb.WriteString(ds[decenas])
		if unidades > 0 {
			sb.WriteString(" y")
		}
	}
	if unidades > 0 {
		sb.WriteRune(' ')
		sb.WriteString(us[unidades])
	} else if n == 0 {
		sb.WriteString(us[0])
	}
	return strings.TrimSpace(sb.String()), err
}

func quotient(numerator, denominator int) int {
	return numerator / denominator
}
