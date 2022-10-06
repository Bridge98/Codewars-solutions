#include <stdbool.h>

bool is_pangram(const char *str_in) {
    char *s; 
    int i, used[26]={0}, total=0;
    
    s = str_in;
    for(i=0;s[i]!='\0';i++)
    {
        if('a'<=s[i] && s[i]<='z')
        {
            total+=!used[s[i]-'a'];
            used[s[i]-'a']=1;
        }
        else if('A'<=s[i] && s[i]<='Z')
        {
            total+=!used[s[i]-'A'];
            used[s[i]-'A']=1;
        }
    }
     
    if(total==26)
    {
        return true;
    }
    else
    {
        return false;
    }

    return true;
}