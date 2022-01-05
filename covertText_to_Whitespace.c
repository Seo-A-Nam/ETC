#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*
*	! convert text-style whitespapce to a whitespace charactor !
*	
*   When to use :
*   	When you get string argument from main(), it has whitespace in text-style.
*       ex) when you print it, then it'll display "hi\tnice to meet you", not "hi   nice to meet you."
*	
*   Cautions :
*	I made this code to convert those text into a new string which i intended to write.
*	Btw, Note that this function make memory allocation through malloc().
*   So don't forget to free() those pointer after you finish to use it. 
*/

static char *handle_notslash(char *str, char *result, int *slash, int *i)
{
	if ((*slash) == 1)  // "was there slash right before this index?"
	{
		(*slash) = 0;
		(*result++) = 92;
	}
	(*result++) = str[*i];
	(*i)++;
	return (result);
}

static char *handle_slash(char *str, char *result, int *slash, int *i)
{
	(*slash) = 0;
	if (str[*i] == 't')
		(*result++) = '\t';
	else if (str[*i] == 'n')
		(*result++) = '\n';
	else if (str[*i] == 'v')
		(*result++) = '\v';
	else if (str[*i] == 'r')
		(*result++) = '\r';
	else if (str[*i] == 'f')
		(*result++) = '\f';
	else if (str[*i] == 92)
	{
		(*result++) = 92;
		(*slash) = 1;
	}
	else
	{
		(*result++) = 92;
		(*result++) = str[*i];
	}
	(*i)++;
	return (result);
}

void    convert_str(char *str, char *result, int len)
{
	int i;
	int slash;

	i = 0;
	slash = 0;
	while (i < len)
	{
		if (str[i] == 92 && slash == 0)
		{
			slash = 1; // "if you find slash, then check it on a variable as a flag."
			i++;
			continue ;
		}
		else if (slash == 1) // "was there slash right before this index?"
			result = handle_slash(str, result, &slash, &i);
		if (str[i] != 92)
			result = handle_notslash(str, result, &slash, &i);
	}
	if (slash == 1) // "was there slash right before this index?"
		(*result++) = str[len - 1];
	*result = '\0';
}

int     get_count(char *str, int len) // Count the number of whitespace charactor existing
{
	int i;
	int count;

	i = 1;
	count = 0;
	while (i < len)
	{
	    if (str[i - 1] == 92)
	    {
            if ((str[i] == 't') || (str[i] == 'n') || (str[i] == 'v') || (str[i] == 'r') || (str[i] == 'f'))
	            count++; 
	    }
	    i++;
	}
	return (count);
}

char    *ft_convert_textws(char *str)
{
	int len;
	int count;
	char *result;

	len = strlen(str);
	count = get_count(str, len);
	if (count == 0)
		return (strdup(str));
	result = (char *)malloc(sizeof(char) * (len - count)); 
	// Whitespace takes 2 slots when it's text style. so minus that count to get a length of new string.
	if (!(result))
		return (NULL);
	convert_str(str, result, len);
	return (result);
} // convert text-style whitespace into a real whitespace charactor. (Make a new string with malloc()).

int     main(int argc, char **argv)
{
    // ex : ./a.out "abc\abc\abc\ta\n\con"
    char *new;
    char *s = argv[1];
    
    new = ft_convert_textws(s);
    printf("argv[1] : %s\n", s);
    printf("result : %s\n", new);
    free(new);
}