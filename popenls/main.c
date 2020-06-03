#include "strings.h"
#include <stdio.h>

int
main()
{
	FILE *f;
	char s[512];

	if((f = popen("ls", "r")) == NULL){
		puts("unable to run ls command");
		return 1;
	}
	while(fgets(s, 512, f)){
		trimnewline(s);
		puts(s);
	}
	pclose(f);
	return 0;
}
