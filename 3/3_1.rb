require 'io/console'

class ThreadPool
  def initialize(size)
    @size = size
    @jobs = Queue.new
    @pool = Array.new(@size) do |i|
      Thread.new do
        Thread.current[:id] = i
        catch(:exit) do
          loop do
            job, args = @jobs.pop
            job.call(*args)
          end
        end
      end
    end
  end

  # add a job to queue
  def schedule(*args, &block)
    @jobs << [block, args]
  end

  # run threads and perform jobs from queue
  def run!
    @size.times do
      schedule { throw :exit }
    end
    @pool.map(&:join)
  end
end

def post(list,flag,name)
  puts "post #{name} started"
  loop do
    qu = gets.chomp
    t = rcislo()
    if list.length >=100
      next
    end
    if list.length<=80
      list.push(t)
      puts "fabric #{name} gave #{t}"
    end
  end
end

def rcislo()
  rc = Random.new
  return rc.rand(100)
end

def consumer(list,flag,name)
  puts "pocup #{name} started"
  loop do
    if list.length >0
      puts "store #{name} get #{list[-1]}"
      list.pop
    end
    if list.length==0 and flag==0
      next
    end
    if list.length==0 and flag>0
      break
    end
  end
end

def quit(flag)
  loop do
    case $stdin.getch
    when 'q'    then exit
    when "\c?"  then puts 'backspace'
    when "\e"   # ANSI escape sequence
      case $stdin.getch
      when '['  # CSI
        case $stdin.getch
        when 'A' then puts 'up'
        when 'B' then puts 'down'
        when 'C' then puts 'right'
        when 'D' then puts 'left'
        end
      end
    end
  end
end


# an instance of ThreadPool with 5 threads
pool = ThreadPool.new(5)
b =[]
flag=0

1.times do
  pool.schedule do
    sleep_time = rand(4) + 2
    sleep(sleep_time)
    quit(flag)
  end
end


2.times do |r|
  pool.schedule do
    sleep_time = rand(4) + 2
    sleep(sleep_time)
    consumer(b,flag,r)
  end
end

# add 3 tasks to query
3.times do |i|
  pool.schedule do
    sleep_time = rand(4) + 2
    sleep(sleep_time)
    post(b,flag,i)
    if i==2
      post(b,flag,i)
    end
  end
end

# run all threads
pool.run!
