<?php

/**/

class MyArray
{
    private $data;
    private $capacity;
    private $length;

    public function __construct($capacity)
    {
        if ($capacity <= 0) {
            return null;
        }
        $this->data = [];
        $this->capacity = $capacity;
        $this->length = 0;
    }

    //判断是否为空
    public function isEmpty()
    {
        if ($this->length == 0) {
            return true;
        }

        return false;
    }

    //判断是否已满
    public function checkIsFull()
    {
        if ($this->length == $this->capacity) {
            return true;
        }

        return false;
    }

    //判断是否越界
    public function checkOutOfRange($index)
    {
        if (($index > (($this->capacity) - 1)) || ($index < 0)) {
            return true;
        } else {
            return false;
        }
    }

    //插入

    /**
     * 最好：O（1）
     * 最坏：O（n）
     * 平均：O（n）
     */
    public function insert($index, $data)
    {
        if ($this->checkIsFull()) {
            echo "FullInsert-----Index:".$index."----data:".$data."\n";
            return 111;
        }

        if ($this->checkOutOfRange($index)) {
            return 222;
        }


        $jump = FALSE;

        //当index跳跃时，补齐
        for ($i = 0;$i<=$index;$i++){
            if(!isset($this->data[$i])){
                $jump = true;
                $this->length++;
                $this->data[$i] = 0;
            }
        }

        if($jump){
            $this->data[$index] = $data;
            return 0;
        }



        //复杂度：O（n）  ，如果是无序的，可以把$index的值搬到最后一个。这样复杂度变为O（1）
        //把index后的元素往后移动一位。
        for ($i = $this->length - 1; $i >= $index; $i--) {
            $this->data[$i + 1] = $this->data[$i];
        }

        $this->data[$index] = $data;
        $this->length++;

        return 0;
    }

    //删除
    public function delete($index)
    {
        if ($this->checkOutOfRange($index)) {
            return 22;
        }

        for ($i = $index; $i < ($this->length - 1); $i++) {
            $this->data[$i] = $this->data[$i + 1];
        }



        $this->length--;

        return 0;
    }

    //find
    public function find($index)
    {
        if ($this->checkOutOfRange($index)) {
            return [
                2,
                0
            ];
        }
        echo $this->data[$index];

        return $this->data[$index];
    }

    //print
    public function printf()
    {
        $format = " ";
        for ($i = 0; $i < $this->length; $i++) {
            $dat = isset($this->data[$i]) ? $this->data[$i] : 0;
            $format .= "|" . $dat;
        }
        echo "\n" . $this->length . "-------" . $format . "\n";
    }
}


$myArr = new MyArray(10);
$myArr->printf();
$myArr->insert(0, 1);
$myArr->printf();
$myArr->insert(2, 2);
$myArr->printf();
$myArr->insert(1, 3);
$myArr->insert(1, 13);
$myArr->delete(1);
$myArr->printf();



$myArr->insert(3, 4);
$myArr->printf();
$myArr->insert(4, 5);
$myArr->insert(5, 6);
$myArr->insert(6, 7);

$myArr->printf();

$myArr->find(3);
$myArr->delete(3);

$myArr->printf();

$myArr->insert(0, 8);
$myArr->insert(4, 9);
$myArr->insert(8, 10);
$myArr->insert(7, 11);
$myArr->printf();