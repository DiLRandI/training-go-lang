using System;
using System.Collections.Generic;

namespace training_demo_csharp
{
    class Program
    {
        static void Main(string[] args)
        {
            Shape square = new Square("Red", 2);
            Shape circle = new Circle("Green", 2);

            AreaCalculateManager mgr = new AreaCalculateManager();
            List<Shape> shapes = new List<Shape>();

            shapes.Add(circle);
            shapes.Add(square);

            var area = mgr.CalculateArea(shapes);
            Console.Write(area.ToString());

            Console.ReadKey();
        }
    }

    public abstract class Shape
    {
        public string Colour { get; }

        protected Shape(string colour)
        {
            Colour = colour;
        }

        public abstract double GetArea();
    }

    public class Square : Shape
    {
        public double Width { get;  }

        public Square(string colour, double width) : base(colour)
        {
            this.Width = width;
        }

        public override double GetArea()
        {
            return Width * Width;
        }
    }

    public class Circle : Shape
    {
        public double Radius { get; }
        public Circle(string colour, double radius) : base(colour)
        {
            Radius = radius;
        }

        public override double GetArea()
        {
            return Math.PI * Math.Pow(Radius, 2);
        }
    }

    public class AreaCalculateManager
    {
        public double CalculateArea(IList<Shape> Shapes)
        {
            double area = 0;
            
            foreach (var shape in Shapes)
            {
                area = area + shape.GetArea();
            }

            return area;
        }
    }
}