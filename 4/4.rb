a = 0
b = 3
c = 3
i = 0
t1=Time.now

until i <= 100000000 do
  a += b * 2 + c - i
  i += 1
end
t2=Time.now

dur=t2-t1
puts dur
