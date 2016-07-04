class A 

def maxDiff(a)
	m = a.max
	i = a.index(m)
	puts i
	sub = a[0...i]
	puts sub
	mm = sub.min
	if mm == nil then
		mm = mm.to_i
	end
	if m != 0 && m != nil then
	 d = m-mm 	
	else 
		d = m
	end
	return d 

end

end
arr = [5, 3]

b =  A.new

puts(b.maxDiff(arr))