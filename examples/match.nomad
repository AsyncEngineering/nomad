num ?
    | 1 => "one"
    | 2 => "two"
    | 3 => "three


item ?
     | str : string => print str
     | num : i32 => print num + 1
     | num : i64 => print num / 2
     | _ => exit(1)
