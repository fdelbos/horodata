angular.module('horodata').factory("searchUsersService", [
   ->

    class Search
      constructor: (@items) ->
        @selected = null
        @searchText = ""
        @index = lunr( ->
          this.field('name')
          this.field('email')
          this.ref('index')
        )
        for i in [0..@items.length - 1]
          obj =
            name: @items[i].name
            email: @items[i].email
            index: i
          @items[i] = obj
          @index.add(obj)

      search: (query) ->
        if query == "" then return @items
        res = @index.search(query)
        data = []
        data.push(@items[i.ref]) for i in res
        return data

      select: (item) ->
        @selected = item

    return {
      get: (items) -> new Search(items)
    }
])
