One suggested idea is to store an array as a JSON object. With json support in sqlite3, this would work pretty well.


This site:
https://www.educba.com/sqlite-array/

Suggested this:

```
Now see how to directly implement arrays in SQLite this is not the right way but till it shows results.

create table aa (stud_id integer, email text [] [], test integer []);

Explanation

In the above example we use create table statement to new table name as aa with different attributes, see here we define email and test as integer array data structure as shown in above statement.

After that insert some records into the demo table by using the following statement as follows.

insert into aa values(‘{{“work”, “rah@gamil.com”}, {“other”, “aaa@gmail.com”}}’, ‘{85, 54, 45}’);
select * from aa;
```

The English is confusing... I'll try this and see if works.
```
$ sqlite3
-- Loading resources from /home/cecil/.sqliterc
SQLite version 3.34.1 2021-01-20 14:10:07
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> create table aa (stud_id integer, email text [] [], test integer []);
sqlite> .table
aa
sqlite> .schema
CREATE TABLE aa (stud_id integer, email text [] [], test integer []);
sqlite> insert into aa values(1,‘{{“work”, “rah@gamil.com”}, {“other”, “aaa@gmail.com”}}’, ‘{85, 54, 45}’);
Error: unrecognized token: "{"
sqlite> 
```
Doesn't work... 
