#include <iostream>
#include <format>

using namespace std;

enum DateType
{
	TYPE_DATE_GREGORIAN,
	TYPE_DATE_JALALI
};

#define MONTHS_COUNT 12
#define DIFF_YEARS 621
#define DIFF_DAYS 81

class Date
{
protected:
	int year, month, day;
	int gMonthDays[MONTHS_COUNT] = {31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29};
	int jMonthDays[MONTHS_COUNT] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
	int *monthDays;
	DateType type;

	void setDateType(DateType t)
	{
		if (t == TYPE_DATE_GREGORIAN)
			monthDays = gMonthDays;
		else if (t == TYPE_DATE_JALALI)
			monthDays = jMonthDays;
		else
			throw invalid_argument("invalid date type");
		type = t;
	}

	void jalaliToGregorian()
	{
		int ny = year, nm = month, nd = day;

		setDateType(TYPE_DATE_GREGORIAN);

		ny += DIFF_YEARS;

		for (int i = 0; i < DIFF_DAYS; i++) {
			nd++;
			if (nd > monthDays[nm - 1]) {
				nm++;
				nd = 1;
			}
			if (nm > MONTHS_COUNT) {
				ny++;
				nm = 1;
			}
		}

		setDate(ny, nm, nd);
	}

	void gregorianToJalali()
	{
		int ny = year, nm = month, nd = day;

		setDateType(TYPE_DATE_JALALI);
		
		ny -= DIFF_YEARS;

		for (int i = 0; i < DIFF_DAYS; i++) {
			nd--;
			if (nd == 0) {
				nm--;
				nd = monthDays[nm - 1];
			}
			if (nm == 0) {
				ny--;
				nm = MONTHS_COUNT;
				nd = monthDays[nm - 1];
			}
		}

		setDate(ny, nm, nd);
	}

public:
	Date() { setDateType(TYPE_DATE_GREGORIAN); }
	Date(DateType t) { setDateType(t); }

	Date(int y, int m, int d)
	{
		setDateType(TYPE_DATE_GREGORIAN);
		setDate(y, m, d);
	}

	Date(DateType t, int y, int m, int d)
	{
		setDateType(t);
		setDate(y, m, d);
	}

	void setDate(int y, int m, int d)
	{
		year = y;
		month = m;
		day = d;
	}

	int getYear() { return year; }
	int getMonth() { return month; }
	int getDay() { return day; }

	string getFormattedDate()
	{
		char date[10];
		sprintf(date, "%d-%d-%d", year, month, day);
		return string(date);
	}

	DateType getType() { return type; }

	Date numberToDate(int n)
	{
		int m, d, days = 0;
		DateType t = getType();

		for (m = 0; m < MONTHS_COUNT; m++)
		{
			if (n <= monthDays[m])
			{
				d = n;
				break;
			}

			n -= monthDays[m];
		}

		return Date(0, m + 1, d);
	}

	void convertTo(DateType t)
	{
		if (t == getType())
			throw invalid_argument("your type is equal current type");

		if (t == TYPE_DATE_GREGORIAN)
			jalaliToGregorian();
		else if (t == TYPE_DATE_JALALI)
			gregorianToJalali();
		else
			throw invalid_argument("invalid date type");
	}
};

int main()
{
	while(1) {
		int o, n, y, m, d;
		char opt;
		DateType t, nt;
		Date date;

		cout << "[1] Number to date" << endl
		     << "[2] Date convertion" << endl
		     << "[3] Exit" << endl
		     << "Chose your option: ";
		cin >> o;
		
		switch (o)
		{
		case 1:
			cout << "Gregorian or jalali (g/j)?: ";
			cin >> opt;
			t = (opt == 'j') ? TYPE_DATE_JALALI : TYPE_DATE_GREGORIAN;

			cout << "Enter your number: ";
			cin >> n;

			cout << Date(t).numberToDate(n).getFormattedDate() << endl;
			break;

		case 2:
			cout << "Say your date type (g/j)?: ";
			cin >> opt;
			if (opt == 'j') {
				t = TYPE_DATE_JALALI;
				nt = TYPE_DATE_GREGORIAN;
			} else {
				t = TYPE_DATE_GREGORIAN;
				nt = TYPE_DATE_JALALI;
			}

			cout << "Enter your date: ";
			scanf("%d-%d-%d", &y, &m, &d);

			date = Date(t, y, m, d);
			date.convertTo(nt);
			cout << date.getFormattedDate() << endl;
			break;

		case 3:
			return (EXIT_SUCCESS);
			break;

		default:
			cout << "Invalid options. Try again!" << endl;
			break;
		}
	}

	return (EXIT_SUCCESS);
}
