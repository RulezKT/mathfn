//some math calculations
package mathfn

import (
	"fmt"
	"math"
)

const (
	SEC_TO_RAD         = 4.8481368110953599359e-6 /* radians per arc second */
	RAD_TO_DEG         = 5.7295779513082320877e1
	PI         float64 = 3.14159265358979323846
	// pi/180 - Ratio for conversion degrees to radians
	// we need it for trigonometric functions
	RAD_RATIO          = 0.017453292519943295
	RAD_PER_ARCSECONDS = 4.8481368110953599359e-6 //STR radians per arc second
	JD2000             = 2451545
	JD1950             = 2433282.5
)

/*
//округляет до 9 цифр после запятой
    public static double round_To_9 (double in_Number) {

        return (double) Math.round(in_Number * 1_000_000_000.0) / 1_000_000_000.0;

    }

    //округляет до places цифр после запятой
    //просто обрезает не округляя ни вверх ни вниз
    public static double round_To(double in_Number, int places) {

        double scale = Math.pow(10, places);

        return (double) (long)(in_Number * scale) / scale;

    }

*/

//https://docs.oracle.com/javase/8/docs/api/java/lang/Math.html#atan2-double-double-
//https://www.geeksforgeeks.org/java-lang-math-atan2-java/
//https://en.wikipedia.org/wiki/Atan2
//стандартные atan2 в Java и Javascript возвращают результат в диапазоне
//-Pi +Pi (-180 +180) градусов
//так как это будет наша Longitude,  необходимо скорректировать
//в диапазон 0..360 градусов

//тесты у стандартной atan2 выиграл, потому что не округляет
// на сегодня самая корректная
func Atn2RAD(y float64, x float64) float64 {

	var phi float64 = 0

	// азимут вектора
	if x == 0 && y == 0 {
		return 0
	}

	var absY float64 = math.Abs(y)
	var absX float64 = math.Abs(x)

	//console.log(`abs_y =  ${abs_y}, abs_x =  ${abs_x},`);

	//https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/atan
	// The Math.atan() method returns a numeric value between - π/2  and  π/2  radians. (+90 and -90 degrees)
	// The angle that the line [(0,0);(x,y)] forms with the x-axis in a Cartesian coordinate system - Math.atan(y / x);
	if absX > absY {
		//arctan
		phi = math.Atan(absY / absX)
	} else {
		//arctan
		phi = PI/2 - math.Atan(absX/absY)
	}

	if x < 0 {
		phi = PI - phi
	}
	if y < 0 {
		phi = -phi
	}

	//console.log(`return value from atn2 =  ${phi}`);
	return phi
}

// на сегодня самая корректная
func Atn2RADWith360Check(y float64, x float64) float64 {

	phi := Atn2RAD(y, x)

	//если меньше 0 или больше 360 градусов, корректируем
	if phi < 0 {
		phi = phi + 2*PI
	}
	if phi > 2*PI {
		phi = phi - 2*PI
	}

	return phi
}

func Atn2RADWith90Check(y float64, x float64) float64 {
	theta := Atn2RAD(y, x)

	//не может быть больше 90 или меньше -90 градусов
	if theta > PI/2 {
		fmt.Printf("\n!!!atn2_RAD_with_90_check has Problems with +-90degrees %05.5f", theta)
	}
	if theta < -PI/2 {
		fmt.Printf("\n!!!atn2_RAD_with_90_check has Problems with +-90degrees %05.5f", theta)
	}

	return theta

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// убирает минус или значения больше 360
// если угол должен быть от 0 до 360
// все в Радианах
func Convert_to_0_360_RAD(longitude float64) float64 {

	coeff := Abs(int(longitude / (2 * PI)))

	if longitude < 0 {
		return longitude + float64(coeff)*2*PI + 2*PI
	} else {
		return longitude - float64(coeff)*2*PI
	}

}

// убирает минус или значения больше 360
// если угол должен быть от 0 до 360
// все в Градусах
func Convert_to_0_360_DEG(longitude float64) float64 {

	coeff := Abs(int(longitude / 360))

	if longitude < 0 {
		return longitude + float64(coeff*360+360)
	} else {
		return longitude - float64(coeff*360)
	}

}
