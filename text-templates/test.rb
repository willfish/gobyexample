require 'erb'

def testing
  <<~EOS
    <!DOCTYPE html>
    <html>
      <head>
        <title><%= @name %></title>
      </head>
      <body>
        <h1><%= @email %></h1>
      </body>
    </html>
  EOS
end

@name = "John Doe"
@email = "John Doe<john.doe@example.com>"
template = ERB.new(testing).result(binding)

puts template


