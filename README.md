# hoo-interview-exercise

this package was created in order to fulfill a few requirements from my interviewer

	the interviewer want a logger that could:
        
		- log to multiple destination
                
			- logger logs to file, udp port and a mocked aws cloud storage
                        
		- support for multiple users
                
			- logger can be initiliazed by passing a  path to the config file
                        
			- different users can give different paths and use the logger
                        
		- limit same logs
                
			- logger checks for previous logs and skips logging if
                        
			   the same message has been logged for n times where n is currently 10.
                           
	NOTES
        
	error handling has been omited in a few places.
        
	aws cloud storage logging is mocked
        
  this repo will be changed to private in a few weeks
  
	FURTHER WORK
        
	make a io.MultiWriter for a default logger? helps if the number of locations increase a lot
