-- create extension age; (need to load this on every session)
load 'age';
-- load '$libdir/plugins/age.so';

-- We recommend adding ag_catalog to your search_path to simplify your queries. The rest of this document will assume you have done so. If you do not, remember to add ‘ag_catalog’ to your cypher query function calls.
SET search_path = ag_catalog, "$user", public;
-- set search_path = ag_catalog

-- create graph
SELECT * FROM ag_catalog.create_graph('test_graph');

-- delete graph
-- SELECT * FROM ag_catalog.drop_graph('test_graph', true);


-- creating vertexes (nodes)
SELECT *
FROM cypher('test_graph', $$
            CREATE (:Person {id: 123, name: 'Aaron', title: 'FE developer', salary: 3200 })
            $$) as (n agtype);
-- 
SELECT *
FROM ag_catalog.cypher('test_graph', $$
            CREATE (:Person {id: 124, name: 'Brandon', title: 'BE developer', salary: 3200 })
            $$) as (n ag_catalog.agtype);
            -- 
SELECT *
FROM ag_catalog.cypher('test_graph', $$
            CREATE (:Person {id: 125, name: 'Charlie', title: 'Fullstack developer', salary: 4000 })
            $$) as (n ag_catalog.agtype);
        -- 
SELECT *
FROM ag_catalog.cypher('test_graph', $$
            CREATE (:Person {id: 126, name: 'Didi', title: 'Web3 Engineer', salary: 6969 })
            $$) as (n ag_catalog.agtype);
            -- 
-- 
SELECT *
FROM ag_catalog.cypher('test_graph', $$
            CREATE (:Person {id: 127, name: 'Delta', title: 'Tech Lead', salary: 10000 })
            $$) as (n ag_catalog.agtype);
-- 
SELECT *
FROM ag_catalog.cypher('test_graph', $$
            CREATE (:Person {id: 128, name: 'Echo', title: 'CTO', salary: 0 })
            $$) as (n ag_catalog.agtype);

-- Return vertexes (nodes)
SELECT *
FROM cypher('test_graph', $$
                       MATCH(v: Person)
                       RETURN v
                       $$) as (v agtype);


-- create edges (connection between nodes)
SELECT *
FROM cypher('test_graph', $$
            MATCH (a: Person), (b: Person)
            WHERE a.name = 'Aaron' AND b.name = 'Brandon'
            CREATE (a)-[e:cowokers]->(b)
            RETURN e
            $$) as (e agtype);

SELECT *
FROM cypher('test_graph', $$
            MATCH (a: Person), (b: Person)
            WHERE a.name = 'Delta' AND b.name = 'Echo'
            CREATE (a)-[e:reports]->(b)
            RETURN e
            $$) as (e agtype);

-- Return a vertex with condition
SELECT *
FROM cypher('test_graph', $$
                       MATCH(v {name: 'Aaron'})
                       RETURN v.title
                       $$) as (v agtype);



-- Two relationship
SELECT *
FROM cypher('test_graph', $$
                       MATCH(v {name: 'Delta'})-[r]-(w{name: 'Echo'})
                       RETURN type(r)
                       $$) as (type agtype);

-- one relationship
SELECT *
FROM cypher('test_graph', $$
                       MATCH(v {name: 'Delta'})-[r]->(w{name: 'Echo'})
                       RETURN type(r)
                       $$) as (type agtype);
                       
                       
-- one relationships                       
SELECT *
FROM cypher('test_graph', $$
                       MATCH(v {name: 'Aaron'})<-[r]-(w{name: 'Charlie'})
                       RETURN type(r)
                       $$) as (type agtype);


-- change a vertex (node)
SELECT *
FROM cypher('test_graph', $$
                       MATCH(v {name: 'Didi'})
                       SET v.title = 'CMO'
                       $$) as (type agtype);



-- check the changed value
SELECT *
FROM cypher('test_graph', $$
                       MATCH(v {name: 'Didi'})
                       RETURN v.title
                       $$) as (type agtype);


-- NORMAL SQL

create table
create table Persons (
  id int,
  name varchar(25),
  city varchar(25),
  hired_year int  
);


INSERT into persons (id,name,city,hired_year)	
values (123, 'Aaron', 'KL', 2017);

INSERT into persons (id,name,city,hired_year)	
values (128, 'Echo', 'KL', 2017);


SELECT * from persons
where name='Echo';


-- hybrid query -> inner query = cypher query
select t.id, t.city, t.hired_year from persons as t where t.name in (
  select name
  from cypher('test_graph', $$
              match(v:Person)
              where v.name = 'Aaron'
              return v.name
              $$) as (name varchar(25))
  );



return names based on IDs
select t.name, t.hired_year from persons as t where t.id in (
  select id
  from cypher('test_graph', $$
              match(v:Person)
              return v.id
              $$) as (id int)
  );
