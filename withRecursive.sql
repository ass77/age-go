create table graph
(node integer not null,
 Edges integer[],
 Primary Key (node));


insert into graph("node", "edges")
values(5, ARRAY[7,8]);

with recursive
search_graph(node, edges, path, level)
as (
  select node, edges, ARRAY[g.node], 1
  from graph g where node = 1
  union all
  select g.node, g.edges,
  path || g.node, level+1
  from graph g, search_graph sg
  where g.node = any(sg.edges)
  and g.node <> all(sg.path)
  ) 
  select * from search_graph
  order by node;
