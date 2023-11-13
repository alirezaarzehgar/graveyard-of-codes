# Entities
## Gregorian Callender
### Properies:
- Count of mount days
- Year
- Month
- Day
### Behavior:
- convert number to date
- Gregorian to jalali


## Jalali Callender
### Properies:
- Count of mount days
- Year
- Month
- Day
### Behavior:
- convert number to date
- Jalali to gregorian

# Design
## Similarities
### Properies:
- Year
- Month
- Day
### Behavior:
- convert number to date

## Difference
- Count of mount days

### Properies:
- Count of mount days
- Jalali to gregorian | Gregorian to jalali

## Needs
- Date class
  - Keep year, month and day
  - convert number to date
  - get/set year, month and day
  - get date format
  - convert date to other dates
  - A flag to detect date type
  - toJalali
  - toGregorian
