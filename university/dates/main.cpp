#include <iostream>
#include <format>

using namespace std;

class Date
{
	int *monthDays;

protected:
	int year, month, day;
	int jMonthDays[12] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
	int gMonthDays[12] = {31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29};

	void setDate(int y, int m, int d)
	{
		if (y <= 0 ||
		    m <= 0 || m > 12 ||
		    d <= 0 || d > monthDays[m - 1])
			throw invalid_argument("setDate: invalid date");

		year = y;
		month = m;
		day = d;
	}

	int getYearCount(int days)
	{
		int r = days % 365;
		return r > 0 ? r : 1;
	}

public:
	Date(bool isJalali)
	{
		monthDays = isJalali ? jMonthDays : gMonthDays;
	}

	Date(bool isJalali, int _year, int _month, int _day)
	{
		monthDays = isJalali ? jMonthDays : gMonthDays;
		setDate(_year, _month, _day);
	}

	Date numberToDate(int n)
	{
		if (n < 1 || n > 365)
			throw invalid_argument("invalid number");

		int y, m, d;
		y = getYearCount(n);

		for (m = 0;; m++)
		{
			if (n <= monthDays[m])
			{
				d = n;
				break;
			}
			n -= monthDays[m];
		}

		return Date(monthDays, y, m + 1, n);
	}

	int getYear() { return year; }

	int getMonth() { return month; }

	int getDay() { return day; }

	string getDate()
	{
		return format("{}-{}-{}", year, month, day);
	}

	int getDays()
	{
		int days = 0;
		for (int i = 0; i <= month; i++)
			days += monthDays[i];
		cout << year << year*365 + days + day << endl;
		return year*365 + days + day;
	}
};

class JalaliDate : public Date
{
public:
	JalaliDate() : Date(true) {}

	JalaliDate(int y, int m, int d) : Date(true, y, m, d) {}

	Date toGregorianDate()
	{
		return Date(false).numberToDate(getDays());
	}
};

class GregorianDate : public Date
{
public:
	GregorianDate() : Date(false) {}

	GregorianDate(int y, int m, int d) : Date(false, y, m, d) {}

	Date toJalali()
	{
		return Date(true).numberToDate(getDays());
	}
};

int main()
{
	JalaliDate d;

	try {
		cout << d.numberToDate(60).getDate() << endl;
		cout << JalaliDate(1402, 2, 2).toGregorianDate().getDate() << endl;
		cout << GregorianDate(2023, 3, 30).toJalali().getDate() << endl;
	} catch (exception &e) {
		cerr << e.what() << endl;
	}

	return (EXIT_SUCCESS);
}
