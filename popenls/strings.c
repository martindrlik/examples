#include "strings.h"

void
trimnewline(char *s)
{
	int i;

	for(i=0; i<512; i++){
		if(s[i]=='\n'){
			s[i] = '\0';
			return;
		}
	}
}
