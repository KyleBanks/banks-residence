# banks-residence
Home automation related projects.

### b-light

My wife and I have a massive 'B' light that contains 12 small bulbs and is powered by two AA batteries. The `b-light` project replaces the battery power and physical switch with a Raspberry Pi running a small [Go HTTP server](./b-light). The [Android](./Android) project contains a [BLightActivity](https://github.com/KyleBanks/banks-residence/blob/master/Android/app/src/main/java/com/kylewbanks/residence/banks/banksresidence/BLightActivity.java) which simply executes the API endpoint running on the Raspberry Pi when a button is clicked. 

