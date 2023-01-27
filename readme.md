
# Booking-TJK 
This is a hotel booking website built with an admin dashboard,a PostgreSQl
database, and an email servise using Mail Hog. All backend is written
in Golang. The website includes a booking calendar and a system for
processing bookings

## Getting Started 🚀  
To get started, clone the repository and install the necessery dependencies.
~~~bash  
  git clone https://github.com/ProninIgorr/booking-tjk.git
  cd booking-tjk
  ./run.sh
~~~

Change flags in run.sh
~~~bash
./booking-tjk -dbname=booking -dbuser=postgres -cache=false
 -production=false -dbpass=
~~~
Download [Mail Hog](https://github.com/mailhog/MailHog) 
and start mail servise at http://localhost:8025.

The website will be available at http://localhost:8080

## Built With 
- Go version 1.19.5
- PostgreSQl
- Uses [Mail Hog](https://github.com/mailhog/MailHog) servise
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [alex edwards SCS](https://github.com/alexedwards/scs) session managment
- Uses [nosurf](https://github.com/justinas/nosurf)


 
## Authors 
- [Igor Pronin](https://www.github.com/ProninIgorr)  


 
## Demo  
[![ALT-TajParadise Demo](http://img.youtube.com/vi/k0g3q5KchhI/0.jpg)](https://www.youtube.com/watch?v=k0g3q5KchhI)