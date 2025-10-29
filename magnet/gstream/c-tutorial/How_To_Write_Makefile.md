* Variables can only be strings. You'll typically want to use __:=__, but __=__ also works.
```sh
files := file1 file2
some_file: $(files)
    echo "Look at this variable: " $(files)
```

* Making multiple targets and you want all of them to run. Make an all target
```sh
all: f1.o f2.o

f1.o f2.o:
    echo $@
```
* *wildcard searches your filesystem for matching filesnames
```sh
thing_wrong := *.o
thing_right := $(wildcard *.o)

all: one two three four
#Fails, because $(thing_wrong) is the string "*.o"
one: $(thing_wrong)
#Stays as *.o if there are no files that match this pattern 
two: *.o
#Works as you would expect. In this case, it does nothing
three: $(thing_right)
#Same as rule three
four: $(wilcard *.o)
```
* Automatic Variables
```sh
hey: one two
    #Outputs "hey", since this is the target name
    echo $@
    #Outputs all prerequisites newer than the target
    echo $?
    #Outputs all prerequisites
    echo $^
    #Outputs the first prerequisite
    echo $<
```
* Pattern rules
```sh
#Define a pattern rule that compiles every .c file into a .o file
%.o : %.c
    $(CC) -c $(CFLAGS) $(CPPFLAGS) $< -o $@
```
* Syntax -targets ...: target-pattern: prereq-patterns...
```sh
objects = foo.o bar.o all.o
$(objects): %.o: %.c
    $(CC) -c $^ -o $@
```
* The __filter__ function is used to select certain elements from a list that match a specifc pattern
```sh
obj_files = foo.result bar.o lose.o
# this select all elements in obj_files that end with .o
filtered_files = $(filter %.o,$(obj_files))
```
* The __call__ function 
```sh
sweet_new_fn = Variable Name: $(0) First: $(1) Second: $(2) Empty Variable: $(3)

all:
	@echo $(call sweet_new_fn, go, tigers)
```
* The __shell__ function
```sh
all: 
	@echo $(shell ls -la)
```
* String substitution


__$(patsubst pattern,replacement,text)__ 

shortcut __$(text:pattern=replacement)__

shortcut __$(text:suffix=replacement)__
```sh
foo := a.o b.o l.a c.o
one := $(patsubst %.o,%.c,$(foo))
# This is a shorthand for the above
two := $(foo:%.o=%.c)
# This is the suffix-only shorthand, and is also equivalent to the above.
three := $(foo:.o=.c)

all:
	echo $(one)
	echo $(two)
	echo $(three)
```
The construct __$(TARGET:=.o)__ in a Makefile utilizes a text substitution reference. This is a powerful feature in Make for transforming lists of filenames or other strings.
